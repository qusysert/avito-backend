package handler

import (
	"context"
	"fmt"
)

// FlushExpiredHandler godoc
//
//	@Summary		Flush expired user_segment entries
//	@Description	delete segment by id; delete all user_segment entries with the given id
//	@Tags			user segment
//	@Success		200	{string}	Status Ok
//	@Router			/flushExpired [delete]
func (h Handler) FlushExpiredHandler(ctx context.Context, _ emptyRequest) (*emptyResponse, error) {
	err := h.service.FlushExpired(ctx)
	if err != nil {
		return nil, fmt.Errorf("error flushing expires: %w", err)
	}
	return &emptyResponse{}, nil
}
