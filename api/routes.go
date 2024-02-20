package api

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
