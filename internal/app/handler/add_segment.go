package handler

import (
	"net/http"
)

// AddSegmentRequest example
type AddSegmentRequest struct {
	Name string `json:"name" example:"AVITO_TEST_SEGMENT" format:"string"`
}

// AddSegmentResponse example
type AddSegmentResponse struct {
	Id int `json:"id" example:"1" format:"integer"`
}

// AddSegmentHandler godoc
//
//	@Summary		Add segment
//	@Description	add segment and get back its' id
//	@Tags			segment
//	@Accept			json
//	@Produce		json
//	@Param 			request body 	AddSegmentRequest true "query params"
//	@Success		200	{object}	AddSegmentResponse
//	@Router			/addSegment [post]
func (h Handler) AddSegmentHandler(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req AddSegmentRequest) (AddSegmentResponse, error) {
		id, err := h.service.AddSegment(r.Context(), req.Name)
		if err != nil {
			return AddSegmentResponse{}, err
		}
		return AddSegmentResponse{Id: id}, err
	})
}
