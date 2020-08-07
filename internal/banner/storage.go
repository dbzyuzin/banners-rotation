package banner

import (
	"context"
)

type Storage interface {
	CreateBanner(ctx context.Context, slot Banner) (int64, error)
	GetAllBanners() ([]Banner, error)
	DeleteBanner(id int64) error
}
