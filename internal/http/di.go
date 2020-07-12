package http

import (
	"github.com/dbzyuzin/banners-rotation.git/internal/storage"
	"net/http"
)

type Di struct {
	Storage *storage.Storage
}

type handlerFunc func(w http.ResponseWriter, r *http.Request, di *Di)
