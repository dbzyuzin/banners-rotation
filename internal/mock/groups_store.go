package mock

import (
	"context"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
)

type SDGroupStoreMock struct {
	CreateFunc       func(ctx context.Context, sdgroup sdgroup.SDGroup) (int64, error)
	GetByIdFunc      func(id int64) (sdgroup.SDGroup, error)
	GetWithLimitFunc func(limit int) ([]sdgroup.SDGroup, error)
	GetAllFunc       func() ([]sdgroup.SDGroup, error)
}

func NewSDGroupStore() *SDGroupStoreMock {
	return &SDGroupStoreMock{}
}

func (s SDGroupStoreMock) Create(ctx context.Context, sdgroup sdgroup.SDGroup) (int64, error) {
	return s.CreateFunc(ctx, sdgroup)
}

func (s SDGroupStoreMock) GetById(id int64) (sdgroup.SDGroup, error) {
	return s.GetByIdFunc(id)
}

func (s SDGroupStoreMock) GetWithLimit(limit int) ([]sdgroup.SDGroup, error) {
	return s.GetWithLimitFunc(limit)
}

func (s SDGroupStoreMock) GetAll() ([]sdgroup.SDGroup, error) {
	return s.GetAllFunc()
}
