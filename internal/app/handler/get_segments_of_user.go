package handler

import (
	"context"
	"fmt"
)

type GetSegmentsOfUserRequest struct {
	Id int
}

type GetSegmentsOfUserResponse struct {
	Segments []Segment `json:"segments"`
}

// GetSegmentsOfUserHandler godoc
//
//	@Summary		Get segments of user
//	@Description	get segments of user by id; returns only not expired entries
//	@Tags			user
//	@Param 			id path int true "User id"
//	@Success		200	{string}	Status Ok
//	@Router			/getSegmentsOfUser/{id} [get]
func (h Handler) GetSegmentsOfUserHandler(ctx context.Context, req GetSegmentsOfUserRequest) (*GetSegmentsOfUserResponse, error) {
	segments, err := h.service.GetSegmentsOfUser(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("cannot get segments of user: %w", err)
	}
	return &GetSegmentsOfUserResponse{Segments: fromModelSegmentList(segments)}, nil
}
