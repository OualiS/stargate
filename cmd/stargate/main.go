package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/fsnotify/fsnotify"
	"log"
	"mrp/config"
	"mrp/router"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

func proxyHandler(cfgAtomic *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := cfgAtomic.Load().(config.Config)
		matchingRoute, isMatchingRoute := router.FindMatchingRoute(cfg, r.Host, r.URL.Path)
		if isMatchingRoute {
			target := fmt.Sprintf("http://%s:%d", matchingRoute.ForwardTo.Container, matchingRoute.ForwardTo.Port)
			parsedUrlTarget, parsedUrlTargetError := url.Parse(target)
			if parsedUrlTargetError != nil {
				http.Error(w, parsedUrlTargetError.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Printf("Matching host detected : %s | Trafic forwarded to : %s\n", matchingRoute.Match.Host, matchingRoute.ForwardTo.Container)
			proxy := httputil.NewSingleHostReverseProxy(parsedUrlTarget)
			proxy.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	}
}

func ReloadConfigIfValid(cfgAtomic *atomic.Value) {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	if err := config.Validate(cfg); err != nil {
		log.Println(err)
		return
	}
	cfgAtomic.Store(cfg)
}

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := config.Validate(cfg); err != nil {
		log.Fatal(err)
	}
	var cfgAtomic atomic.Value
	cfgAtomic.Store(cfg)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	fmt.Println("-------------------------")
	fmt.Println("|  MRP - Reverse proxy  |")
	fmt.Println("|    made by Salah      |")
	fmt.Println("-------------------------")
	fmt.Printf("%d configurations founded\n", len(cfg.Routes))

	go func() {
		fmt.Println("Configuration watcher started")
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Println("File ", event.Name, " has been updated")
					ReloadConfigIfValid(&cfgAtomic)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", proxyHandler(&cfgAtomic))
	myFigure := figure.NewColorFigure("MRP", "", "green", true)
	myFigure.Print()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
