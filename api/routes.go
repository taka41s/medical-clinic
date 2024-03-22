package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"ajp-medical-clinic/api/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler http.Handler // Change type to http.Handler
}

var routes = []Route{
	{Path: "/user", Method: http.MethodPost, Handler: http.HandlerFunc(handlers.RegisterUser)},
	{Path: "/user/{id}", Method: http.MethodGet, Handler: http.HandlerFunc(handlers.FetchUserByID)},
	{Path: "/users", Method: http.MethodGet, Handler: http.HandlerFunc(handlers.FetchUsers)},
	{Path: "/health_insurances", Method: http.MethodGet, Handler: http.HandlerFunc(handlers.FetchHealthInsurances)},
}

func InitializeRoutes() {
	r := mux.NewRouter()

	for _, route := range routes {
		r.HandleFunc(route.Path, wrapHandler(route.Handler)).Methods(route.Method)
	}

	http.Handle("/", r)
}

func wrapHandler(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}
