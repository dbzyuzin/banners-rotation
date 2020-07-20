package mock

import (
	"context"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
)

type SDGroupStoreMock struct {
	CreateFunc func(ctx context.Context, sdgroup sdgroup.SDGroup) (int64, error)
	DeleteFunc func(id int64) error
	GetAllFunc func() ([]sdgroup.SDGroup, error)
}

func NewSDGroupStore() *SDGroupStoreMock {
	return &SDGroupStoreMock{}
}

func (s SDGroupStoreMock) Create(ctx context.Context, sdgroup sdgroup.SDGroup) (int64, error) {
	return s.CreateFunc(ctx, sdgroup)
}

func (s SDGroupStoreMock) Delete(id int64) error {
	return s.DeleteFunc(id)
}

func (s SDGroupStoreMock) GetAll() ([]sdgroup.SDGroup, error) {
	return s.GetAllFunc()
}
