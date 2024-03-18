package api

import (
	"net/http"
	"strings"
	"ajp-medical-clinic/api/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler http.Handler // Change type to http.Handler
}

var routes = []Route{
	{Path: "/user", Method: http.MethodPost, Handler: http.HandlerFunc(handlers.RegisterUser)},
	{Path: "/users", Method: http.MethodGet, Handler: http.HandlerFunc(handlers.FetchUsers)},
}

func InitializeRoutes() {
	for _, route := range routes {
		http.Handle(route.Path, MethodMiddleware(route.Handler))
	}
}

func MethodMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsMethodAllowed(r.URL.Path, r.Method) {
			http.Error(w, "BAD REQUEST", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsMethodAllowed(path string, method string) bool {
	for _, route := range routes {
		if route.Path == path {
			return strings.ToUpper(route.Method) == strings.ToUpper(method)
		}
	}
	return false
}