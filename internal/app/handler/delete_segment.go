package handler

import (
	"context"
	"fmt"
)

type DeleteSegmentRequest struct {
	Name string `query:"name"`
}

// DeleteSegmentHandler godoc
//
//	@Summary		Delete segment
//	@Description	delete segment by id; delete all user_segment entries with the given id
//	@Tags			segment
//	@Param 			id path int true "Segment id"
//	@Success		200	{string}	Status Ok
//	@Router			/deleteSegment/{id} [delete]
func (h Handler) DeleteSegmentHandler(ctx context.Context, req DeleteSegmentRequest) (*emptyResponse, error) {
	err := h.service.DeleteSegment(ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("cannot delete segment: %w", err)
	}
	return &emptyResponse{}, err
}
