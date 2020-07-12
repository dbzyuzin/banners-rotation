package main

import (
	handler "github.com/dbzyuzin/banners-rotation.git/internal/http"
	"log"
	"net/http"
)

func main() {

	h, _ := handler.NewHandler()

	log.Fatal(http.ListenAndServe(":80", h))
}
