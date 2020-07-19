package http

import (
	"github.com/dbzyuzin/banners-rotation.git/internal/mock"
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

func TestCreate1(t *testing.T) {
	h, _ := NewHandler(&storage.Storage{SDGroups: mock.NewSDGroupStore()})
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
