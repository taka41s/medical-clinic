package api

import (
	"net/http"
	"ajp-medical-clinic/api/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	{Path: "/user", Method: http.MethodPost, Handler: handlers.RegisterUser},
	{Path: "/users", Method: http.MethodGet, Handler: handlers.FetchUsers},
}
