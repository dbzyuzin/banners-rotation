package memstore

import "github.com/dbzyuzin/banners-rotation.git/internal/storage"

func New() *storage.Storage {
	return &storage.Storage{SDGroups: NewSDGroupStore()}
}
