package slot

import "context"

type Storage interface {
	CreateSlot(ctx context.Context, slot Slot) (int64, error)
	GetAllSlots() ([]Slot, error)
	DeleteSlot(id int64) error
}
