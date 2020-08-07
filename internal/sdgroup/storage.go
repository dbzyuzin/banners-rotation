package sdgroup

import (
	"context"
)

type Storage interface {
	CreateGroup(ctx context.Context, sdgroup SDGroup) (int64, error)
	GetAllGroups() ([]SDGroup, error)
	DeleteGroup(id int64) error
}
