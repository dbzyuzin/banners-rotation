package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler() (http.Handler, error) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	router.HandleFunc("/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		w.Write([]byte("Hello, " + name))
	})

	return router, nil
}
