package http

import (
	"github.com/dbzyuzin/banners-rotation.git/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(store *storage.Storage) (http.Handler, error) {
	router := mux.NewRouter()

	groupsHandler := NewSDGroupHandler(store.SDGroups)

	groups := router.PathPrefix("/sd-groups").Subrouter()
	groups.Handle("", asHandler(groupsHandler.Create)).Methods("POST")
	groups.Handle("", asHandler(groupsHandler.GetAll)).Methods("GET")

	return router, nil
}

func asHandler(in func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return in
}
