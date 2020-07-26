package banner

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, slot Banner) (int64, error)
	GetAll() ([]Banner, error)
	Delete(id int64) error
}
