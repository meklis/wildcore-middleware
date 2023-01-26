package wca_auth

import (
	"net/http"
)

const PluginName = "wca_auth"

type Plugin struct{}

// to declare plugin
func (p *Plugin) Init() error {
	return nil
}

func (p *Plugin) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something
		// ...
		// continue request through the middleware pipeline
		next.ServeHTTP(w, r)
	})
}

// Middleware/plugin name.
func (p *Plugin) Name() string {
	return PluginName
}
