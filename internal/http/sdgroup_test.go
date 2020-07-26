package http

import (
	"encoding/json"
	"github.com/dbzyuzin/banners-rotation.git/internal/mock"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
	"github.com/dbzyuzin/banners-rotation.git/internal/storage"
	"github.com/dbzyuzin/banners-rotation.git/internal/storage/memstore"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	h, _ := NewHandler(memstore.New())
	server := httptest.NewServer(h)
	defer server.Close()

	resp, err := http.Post(server.URL+"/sd-groups", "text/plain", strings.NewReader("it is description"))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	r, _ := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	_, err = strconv.Atoi(string(r))
	require.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	groupStore := mock.NewSDGroupStore()
	h, _ := NewHandler(&storage.Storage{SDGroups: groupStore})
	server := httptest.NewServer(h)
	defer server.Close()

	groupStore.GetAllFunc = func() ([]sdgroup.SDGroup, error) {
		return []sdgroup.SDGroup{{Id: 1, Description: "text"}}, nil
	}

	resp, err := http.Get(server.URL + "/sd-groups")
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
