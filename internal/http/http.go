package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(di *Di) (http.Handler, error) {
	router := mux.NewRouter()

	monkey := func(handler handlerFunc) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, di)
		}
	}

	router.HandleFunc("/", monkey(hello))

	return router, nil
}

var hello = handlerFunc(func(w http.ResponseWriter, r *http.Request, di *Di) {
	w.Write([]byte("Text"))
})
