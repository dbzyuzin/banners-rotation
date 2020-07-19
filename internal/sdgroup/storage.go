package sdgroup

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, sdgroup SDGroup) (int64, error)
	GetAll() ([]SDGroup, error)
	Delete(id int) error
}
