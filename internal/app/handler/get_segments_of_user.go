package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetSegmentsOfUserHandler godoc
//
//	@Summary		Get segments of user
//	@Description	get segments of user by id; returns only not expired entries
//	@Tags			user
//	@Param 			id path int true "User id"
//	@Success		200	{string}	Status Ok
//	@Router			/getSegmentsOfUser/{id} [get]
func (h Handler) GetSegmentsOfUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idArg := vars["id"]
	id, err := strconv.ParseInt(idArg, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad Request: %v", err), http.StatusBadRequest)
		return
	}
	segments, err := h.service.GetSegmentsOfUser(r.Context(), int(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}

	if len(segments) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	segmentsJSON, err := json.Marshal(segments)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(segmentsJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}
}
