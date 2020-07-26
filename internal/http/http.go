package http

import (
	"github.com/dbzyuzin/banners-rotation.git/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(store *storage.Storage) (http.Handler, error) {
	router := mux.NewRouter()

	groupsHandler := NewSDGroupHandler(store.SDGroups)
	bannersHandler := NewBannerHandler(store.Banners)
	slotsHandler := NewSlotHandler(store.Slots)

	groups := router.PathPrefix("/sd-groups").Subrouter()
	groups.Handle("", http.HandlerFunc(groupsHandler.Create)).Methods("POST")
	groups.Handle("", http.HandlerFunc(groupsHandler.GetAll)).Methods("GET")
	groups.Handle("", http.HandlerFunc(groupsHandler.Delete)).Methods("DELETE")

	banners := router.PathPrefix("/banners").Subrouter()
	banners.Handle("", http.HandlerFunc(bannersHandler.Create)).Methods("POST")
	banners.Handle("", http.HandlerFunc(bannersHandler.GetAll)).Methods("GET")
	banners.Handle("", http.HandlerFunc(bannersHandler.Delete)).Methods("DELETE")

	slots := router.PathPrefix("/slots").Subrouter()
	slots.Handle("", http.HandlerFunc(slotsHandler.Create)).Methods("POST")
	slots.Handle("", http.HandlerFunc(slotsHandler.GetAll)).Methods("GET")
	slots.Handle("", http.HandlerFunc(slotsHandler.Delete)).Methods("DELETE")

	return router, nil
}
