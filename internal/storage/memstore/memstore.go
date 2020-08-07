package memstore

import (
	"context"
	"github.com/dbzyuzin/banners-rotation.git/internal/banner"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
	"github.com/dbzyuzin/banners-rotation.git/internal/slot"
	"github.com/dbzyuzin/banners-rotation.git/internal/storage"
	"sync"
)

func New() storage.Storage {
	return &Store{
		gmux: &sync.Mutex{},
		smux: &sync.Mutex{},
		bmux: &sync.Mutex{},
	}
}

type Store struct {
	gmux         *sync.Mutex
	lastIdGroups int64
	groups       []sdgroup.SDGroup

	bmux          *sync.Mutex
	lastIdBanners int64
	banners       []sdgroup.SDGroup

	smux        *sync.Mutex
	lastIdSlots int64
	slots       []sdgroup.SDGroup
}

func (s *Store) CreateBanner(ctx context.Context, slot banner.Banner) (int64, error) {
	panic("implement me")
}

func (s *Store) GetAllBanners() ([]banner.Banner, error) {
	panic("implement me")
}

func (s *Store) DeleteBanner(id int64) error {
	panic("implement me")
}

func (s *Store) CreateSlot(ctx context.Context, slot slot.Slot) (int64, error) {
	panic("implement me")
}

func (s *Store) GetAllSlots() ([]slot.Slot, error) {
	panic("implement me")
}

func (s *Store) DeleteSlot(id int64) error {
	panic("implement me")
}

func (s *Store) DeleteGroup(id int64) error {
	s.gmux.Lock()
	defer s.gmux.Unlock()
	for i, group := range s.groups {
		if group.Id == id {
			s.groups = append(s.groups[:i], s.groups[i+1:]...)
			return nil
		}
	}
	return sdgroup.ErrNotFound
}

func (s *Store) CreateGroup(_ context.Context, sdgroup sdgroup.SDGroup) (int64, error) {
	s.gmux.Lock()
	defer s.gmux.Unlock()
	s.lastIdGroups++
	sdgroup.Id = s.lastIdGroups
	s.groups = append(s.groups, sdgroup)
	return sdgroup.Id, nil
}

func (s Store) GetAllGroups() ([]sdgroup.SDGroup, error) {
	s.gmux.Lock()
	defer s.gmux.Unlock()

	result := make([]sdgroup.SDGroup, 0, len(s.groups))
	for _, group := range s.groups {
		result = append(result, group)
	}

	return result, nil
}
