package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DeleteSegmentHandler godoc
//
//	@Summary		Delete segment
//	@Description	delete segment by id; delete all user_segment entries with the given id
//	@Tags			segment
//	@Param 			id path int true "Segment id"
//	@Success		200	{string}	Status Ok
//	@Router			/deleteSegment/{id} [delete]
func (h Handler) DeleteSegmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idArg := vars["id"]
	id, err := strconv.ParseInt(idArg, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad Request: %v", err), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteSegment(r.Context(), int(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}
	http.Error(w, fmt.Sprintf("Status Ok"), http.StatusOK)
	return
}
