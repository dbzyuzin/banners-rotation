package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dbzyuzin/banners-rotation.git/internal/slot"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SlotHandler struct {
	Store slot.Storage
}

func NewSlotHandler(store slot.Storage) *SlotHandler {
	return &SlotHandler{
		store,
	}
}

func (s SlotHandler) Create(w http.ResponseWriter, r *http.Request) {
	bodyReader := io.LimitReader(r.Body, 1<<20)
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	description := string(body)

	id, err := s.Store.Create(context.Background(), slot.Slot{Description: description})
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

func (s SlotHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	values, err := s.Store.GetAll()
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

func (s SlotHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.Store.Delete(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
