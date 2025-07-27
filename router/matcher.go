package router

import (
	"mrp/config"
	"strings"
)

func FindMatchingRoute(cfg config.Config, host string, urlPath string) (*config.Route, bool) {
	for i := range cfg.Routes {
		route := &cfg.Routes[i]
		if route.Match.Host == host && strings.HasPrefix(urlPath, route.Match.PathPrefix) {
			return route, true
		}
	}
	return nil, false
}
