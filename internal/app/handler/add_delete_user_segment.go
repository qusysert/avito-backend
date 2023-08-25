package handler

import (
	"avito-backend/internal/app/model"
	"net/http"
)

// AddDeleteUserSegmentRequest example
type AddDeleteUserSegmentRequest struct {
	UserId   int             `json:"user_id" example:"1000" format:"integer"`
	ToAdd    []model.Segment `json:"to_add"`
	ToDelete []string        `json:"to_delete" example:"SEGMENT1,SEGMENT2" format:"array"`
}

// AddDeleteUserSegmentResponse example
type AddDeleteUserSegmentResponse struct {
	AddedIds []int `json:"added_ids" example:"6,7,8" format:"array"`
}

// AddDeleteUserSegmentHandler godoc
//
//	@Summary		Add and delete segments from user by id
//	@Description	add and delete segments from user by id
//	@Tags			user segment
//	@Accept			json
//	@Produce		json
//	@Param request  body AddDeleteUserSegmentRequest true "query params"
//	@Success		200	{object}	AddDeleteUserSegmentResponse
//	@Router			/addDeleteUserSegment [put]
func (h Handler) AddDeleteUserSegmentHandler(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req AddDeleteUserSegmentRequest) (AddDeleteUserSegmentResponse, error) {
		addedIds, err := h.service.AddDeleteUserSegment(r.Context(), req.UserId, req.ToAdd, req.ToDelete)
		if err != nil {
			return AddDeleteUserSegmentResponse{}, err
		}
		return AddDeleteUserSegmentResponse{AddedIds: addedIds}, nil
	})
}
