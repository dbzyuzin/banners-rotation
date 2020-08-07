package memstore

import (
	"context"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAndGet(t *testing.T) {
	store := New()

	description := "some text"
	id, err := store.SDGroups.CreateGroup(context.Background(), sdgroup.SDGroup{
		Description: description,
	})
	require.NoError(t, err)

	values, err := store.SDGroups.GetAllGroups()
	require.NoError(t, err)
	require.Len(t, values, 1)

	value := values[0]
	require.Equal(t, id, value.Id)
	require.Equal(t, description, value.Description)
}
