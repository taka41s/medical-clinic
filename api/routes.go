package api

import 	( "net/http"
		  "ajp-medical-clinic/api/handlers"
		)

type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	{Path: "/user", Handler: handlers.RegisterUser},
	{Path: "/users", Handler: handlers.FetchUsers},
}
