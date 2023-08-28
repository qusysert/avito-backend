package service

import (
	"avito-backend/internal/app/model"
	"avito-backend/internal/pkg/config"
	"avito-backend/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_AddDeleteUserSegment(t *testing.T) {
	repo := mocks.NewIRepository(t)
	now := time.Now()
	ctx := context.Background()

	userSegmentId := 99
	segmentId := 1
	userId := 10
	nonExistingSegmentName := "NON_EXISTING_SEGMENT"
	existingSegmentName := "SEGMENT_1"

	repo.On("AddSegmentIfNotExists", ctx, existingSegmentName).Return(segmentId, nil)
	repo.On("AddUserSegmentIfNotExists", ctx, userId, segmentId, &now).Return(userSegmentId, nil)
	repo.On("DeleteUserSegmentIfExists", ctx, userId, nonExistingSegmentName).Return(false, nil)
	repo.On("DeleteUserSegmentIfExists", ctx, userId, existingSegmentName).Return(true, nil)

	srv := New(config.Config{}, repo)

	for _, c := range []struct {
		userId               int
		toAdd                []model.SegmentWithExpires
		toDelete             []string
		expectedErrorContain string
		expectedValue        []int
	}{
		{
			// Adding segment
			userId:        userId,
			toAdd:         []model.SegmentWithExpires{{Name: existingSegmentName, Expires: now}},
			expectedValue: []int{userSegmentId}},

		{
			// Deleting existing segment
			userId:   userId,
			toDelete: []string{existingSegmentName}},
		{
			// Deleting non-existing segment
			userId:               userId,
			toDelete:             []string{nonExistingSegmentName},
			expectedErrorContain: "no such segment to delete"},
	} {

		actual, err := srv.AddDeleteUserSegment(ctx, c.userId, c.toAdd, c.toDelete)
		if c.expectedErrorContain != "" {
			assert.Error(t, err)
			assert.Contains(t, err.Error(), c.expectedErrorContain)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, c.expectedValue, actual)
	}
}
