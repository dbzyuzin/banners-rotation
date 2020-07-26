package mock

import (
	"context"
	"github.com/dbzyuzin/banners-rotation.git/internal/slot"
)

type SlotStore struct {
	CreateFunc func(ctx context.Context, slot slot.Slot) (int64, error)
	DeleteFunc func(id int64) error
	GetAllFunc func() ([]slot.Slot, error)
}

func NewSlotStore() *SDGroupStoreMock {
	return &SDGroupStoreMock{}
}

func (s SlotStore) Create(ctx context.Context, slot slot.Slot) (int64, error) {
	return s.CreateFunc(ctx, slot)
}

func (s SlotStore) Delete(id int64) error {
	return s.DeleteFunc(id)
}

func (s SlotStore) GetAll() ([]slot.Slot, error) {
	return s.GetAllFunc()
}
