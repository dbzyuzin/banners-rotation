package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dbzyuzin/banners-rotation.git/internal/sdgroup"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SDGroupHandler struct {
	Store sdgroup.Storage
}

func NewSDGroupHandler(store sdgroup.Storage) *SDGroupHandler {
	return &SDGroupHandler{
		store,
	}
}

func (s SDGroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	bodyReader := io.LimitReader(r.Body, 1<<20)
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	description := string(body)

	id, err := s.Store.CreateGroup(context.Background(), sdgroup.SDGroup{Description: description})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte(fmt.Sprint(id)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s SDGroupHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	values, err := s.Store.GetAllGroups()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(&values)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s SDGroupHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		if err == sdgroup.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.Store.DeleteGroup(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
