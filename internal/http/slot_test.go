package http

import (
	"context"
	"encoding/json"
	"github.com/dbzyuzin/banners-rotation.git/internal/mock"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
	"github.com/dbzyuzin/banners-rotation.git/internal/slot"
	"github.com/dbzyuzin/banners-rotation.git/internal/storage"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestCreateSlot(t *testing.T) {
	slotsStoreMock := mock.NewSlotStore()
	store := &storage.Storage{Slots: slotsStoreMock}
	h, err := NewHandler(store)
	require.NoError(t, err)
	server := httptest.NewServer(h)
	defer server.Close()

	t.Run("post for create, store used at least ones", func(t *testing.T) {
		used := 0
		slotsStoreMock.CreateFunc = func(ctx context.Context, slot slot.Slot) (int64, error) {
			used++
			return 0, nil
		}

		_, err := http.Post(server.URL+"/slots", "text/plain", strings.NewReader("it is description"))
		require.NoError(t, err)
		require.Equal(t, used, 1)
	})

	resp, err := http.Post(server.URL+"/slots", "text/plain", strings.NewReader("it is description"))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	r, _ := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	_, err = strconv.Atoi(string(r))
	require.NoError(t, err)
}

func TestGetAllSlots(t *testing.T) {
	groupStore := mock.NewSDGroupStore()
	h, _ := NewHandler(&storage.Storage{SDGroups: groupStore})
	server := httptest.NewServer(h)
	defer server.Close()

	groupStore.GetAllFunc = func() ([]sdgroup.SDGroup, error) {
		return []sdgroup.SDGroup{{Id: 1, Description: "text"}}, nil
	}

	resp, err := http.Get(server.URL + "/slots")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var res []sdgroup.SDGroup
	data, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	err = json.Unmarshal(data, &res)
	require.NoError(t, err)

	require.Len(t, res, 1)
	require.Equal(t, int64(1), res[0].Id)
}
