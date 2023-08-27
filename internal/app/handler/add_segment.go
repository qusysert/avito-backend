package handler

import (
	"context"
	"fmt"
)

// AddSegmentRequest example
type AddSegmentRequest struct {
	Name string `json:"name" validate:"required" example:"AVITO_TEST_SEGMENT" format:"string"`
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
func (h Handler) AddSegmentHandler(ctx context.Context, req AddSegmentRequest) (*AddSegmentResponse, error) {
	id, err := h.service.AddSegment(ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("cannot add segment: %w", err)
	}
	return &AddSegmentResponse{Id: id}, nil
}
