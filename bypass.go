package bypass

import "net/http"

// Bypass is a type that maps paths to their respective handlers.
type Bypass map[string]http.HandlerFunc

// NewBypass is a convenience function to return a new Bypass map.
func NewBypass() Bypass {
	return make(Bypass)
}

// AddPath adds a path with an associated handler. If you add a path more than once it will
// overwrite the existing handler for that path.
func (p Bypass) AddPath(path string, h http.HandlerFunc) {
	p[path] = h
}

// Handle is the middleware for Bypass. It checks if the path in the Request matches any
// in the map and if so calls its function. This allows you to skip middleware chains
// for certain routes.
func (p Bypass) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if pass, ok := p[r.URL.Path]; ok {
			pass(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
