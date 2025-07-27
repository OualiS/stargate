package router

import (
	"mrp/config"
	"strings"
)

// FindMatchingRoute searches for a route in the provided configuration that matches the given host and URL path prefix.
// Returns a pointer to the matching route and a boolean indicating whether a match was found.
func FindMatchingRoute(cfg config.Config, host string, urlPath string) (*config.Route, bool) {
	for i := range cfg.Routes {
		route := &cfg.Routes[i]
		if route.Match.Host == host && strings.HasPrefix(urlPath, route.Match.PathPrefix) {
			return route, true
		}
	}
	return nil, false
}
