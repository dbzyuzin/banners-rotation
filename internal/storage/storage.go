package storage

import (
	"github.com/dbzyuzin/banners-rotation.git/internal/banner"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
	"github.com/dbzyuzin/banners-rotation.git/internal/slot"
)

type Storage interface {
	sdgroup.Storage
	banner.Storage
	slot.Storage
}
