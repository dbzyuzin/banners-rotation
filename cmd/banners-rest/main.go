package main

import (
	handler "github.com/dbzyuzin/banners-rotation.git/internal/http"
	"github.com/dbzyuzin/banners-rotation.git/internal/storage/memstore"
	"log"
	"net/http"
)

func main() {
	h, _ := handler.NewHandler(memstore.New())

	log.Fatal(http.ListenAndServe(":80", h))
}
