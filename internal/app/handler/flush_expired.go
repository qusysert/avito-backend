package handler

import (
	"fmt"
	"net/http"
)

// FlushExpiredHandler godoc
//
//	@Summary		Flush expired user_segment entries
//	@Description	delete segment by id; delete all user_segment entries with the given id
//	@Tags			user segment
//	@Success		200	{string}	Status Ok
//	@Router			/flushExpired [delete]
func (h Handler) FlushExpiredHandler(w http.ResponseWriter, r *http.Request) {
	err := h.service.FlushExpired(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}
	http.Error(w, fmt.Sprintf("Status Ok"), http.StatusOK)
}
