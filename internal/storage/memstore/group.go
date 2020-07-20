package memstore

import (
	"context"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
	"sync"
)

type SDGroupStore struct {
	mux    *sync.Mutex
	lastId int64
	groups []sdgroup.SDGroup
}

func NewSDGroupStore() *SDGroupStore {
	return &SDGroupStore{
		mux: &sync.Mutex{},
	}
}

func (s *SDGroupStore) Create(ctx context.Context, sdgroup sdgroup.SDGroup) (int64, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.lastId++
	sdgroup.Id = s.lastId
	s.groups = append(s.groups, sdgroup)
	return sdgroup.Id, nil
}

func (s SDGroupStore) GetAll() ([]sdgroup.SDGroup, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	result := make([]sdgroup.SDGroup, 0, len(s.groups))
	for _, group := range s.groups {
		result = append(result, group)
	}

	return result, nil
}
