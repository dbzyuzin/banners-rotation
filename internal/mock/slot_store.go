package mock

import (
	"context"
	"github.com/dbzyuzin/banners-rotation.git/internal/slot"
)

type SlotStoreMock struct {
	CreateFunc func(ctx context.Context, slot slot.Slot) (int64, error)
	DeleteFunc func(id int64) error
	GetAllFunc func() ([]slot.Slot, error)
}

func NewSlotStore() *SlotStoreMock {
	return &SlotStoreMock{}
}

func (s SlotStoreMock) CreateSlot(ctx context.Context, slot slot.Slot) (int64, error) {
	return s.CreateFunc(ctx, slot)
}

func (s SlotStoreMock) DeleteSlot(id int64) error {
	return s.DeleteFunc(id)
}

func (s SlotStoreMock) GetAllSlots() ([]slot.Slot, error) {
	return s.GetAllFunc()
}
