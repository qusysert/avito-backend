package handler

import (
	"avito-backend/internal/app/model"
	"fmt"
	"time"
)

type Segment struct {
	Name    string `json:"name" validate:"required" example:"SEGMENT10"`
	Expires string `json:"expires" validate:"datetime" example:"2023-08-25T17:00:05"`
}

func fromModelSegment(m model.SegmentWithExpires) Segment {
	return Segment{
		Name:    m.Name,
		Expires: m.Expires.Format("2006-01-02T15:04:05"),
	}
}

func toModelHandlerSlice(segments []Segment) ([]model.SegmentWithExpires, error) {
	var modelSegments []model.SegmentWithExpires

	for _, m := range segments {
		expires, err := time.Parse("2006-01-02T15:04:05", m.Expires)
		if err != nil {
			return nil, fmt.Errorf("cannot parse expires: %v", err)
		}

		modelSegment := model.SegmentWithExpires{
			Name:    m.Name,
			Expires: expires,
		}

		modelSegments = append(modelSegments, modelSegment)
	}

	return modelSegments, nil
}

func fromModelSegmentList(ml []model.SegmentWithExpires) []Segment {
	ret := make([]Segment, 0, len(ml))
	for _, s := range ml {
		ret = append(ret, fromModelSegment(s))
	}
	return ret
}
