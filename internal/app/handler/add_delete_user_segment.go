package handler

import (
	"context"
	"fmt"
)

// AddDeleteUserSegmentRequest example
type AddDeleteUserSegmentRequest struct {
	UserId   int       `json:"user_id" validate:"required" example:"1000" format:"integer"`
	ToAdd    []Segment `json:"to_add"`
	ToDelete []string  `json:"to_delete" example:"SEGMENT1,SEGMENT2" format:"array"`
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
func (h Handler) AddDeleteUserSegmentHandler(ctx context.Context, req AddDeleteUserSegmentRequest) (*AddDeleteUserSegmentResponse, error) {
	toAdd, err := toModelHandlerSlice(req.ToAdd)
	if err != nil {
		return nil, fmt.Errorf("troubles with converting structure: %v", err)
	}
	addedIds, err := h.service.AddDeleteUserSegment(ctx, req.UserId, toAdd, req.ToDelete)
	if err != nil {
		return nil, fmt.Errorf("cannot add and delete segmets from user: %w", err)
	}
	return &AddDeleteUserSegmentResponse{AddedIds: addedIds}, nil
}
