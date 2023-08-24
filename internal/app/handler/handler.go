package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"net/http"
	"reflect"
)

type IService interface{}

type Handler struct {
	service IService
}

func New(srv IService) *Handler {
	return &Handler{service: srv}
}

// Unmarshal request, do work fn(), then marshall response into JSON anf return
func handle[REQ any, RESP any](w http.ResponseWriter, r *http.Request, fn func(req REQ) (RESP, error)) {
	var req REQ
	body, err := io.ReadAll(r.Body)
	if err != nil {
		toWriterError(&w, err, http.StatusBadRequest)
		return
	}

	if !isNil(req) {
		err = json.Unmarshal(body, &req)
		if err != nil {
			toWriterError(&w, err, http.StatusBadRequest)
			return
		}
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		toWriterError(&w, err, http.StatusBadRequest)
		return
	}

	resp, err := fn(req)
	if err != nil {
		toWriterError(&w, err, http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		toWriterError(&w, err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respJSON)
	if err != nil {
		toWriterError(&w, err, http.StatusInternalServerError)
		return
	}
}

func toWriterError(w *http.ResponseWriter, er error, code int) {
	http.Error(*w, fmt.Sprintf("%+v", er), code)
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
