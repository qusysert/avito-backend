package handler

import (
	"context"
	"fmt"
)

type DeleteSegmentRequest struct {
	Name string `query:"name" validate:"required"`
}

// DeleteSegmentHandler godoc
//
// @Tags segment
// @Summary Delete segment
// @Description delete segment by name; delete segment and all user_segment entries with the given segment name
// @Param name query string true "Segment name"
// @Success 200 {string} Status Ok
// @Router /deleteSegment [delete]
func (h Handler) DeleteSegmentHandler(ctx context.Context, req DeleteSegmentRequest) (*emptyResponse, error) {
	err := h.service.DeleteSegment(ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("cannot delete segment: %w", err)
	}
	return &emptyResponse{}, err
}
