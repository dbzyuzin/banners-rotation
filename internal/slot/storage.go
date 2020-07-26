package slot

import "context"

type Storage interface {
	Create(ctx context.Context, slot Slot) (int64, error)
	GetAll() ([]Slot, error)
	Delete(id int64) error
}
