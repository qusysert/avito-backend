package test

import (
	"avito-backend/internal/app/handler"
	"math/rand"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	env, rollback := NewEnv()
	defer rollback()

	segments := make([]string, 4)
	expires := make([]string, 4)
	userIds := make([]int, 4)

	for i := 0; i < 4; i++ {
		segments[i] = genName()
		expires[i] = genDate()
		userIds[i] = rand.Intn(10000) + 1
	}

	respAddSegment, err := env.Hdl.AddSegmentHandler(env.Ctx, handler.AddSegmentRequest{Name: segments[0]})
	if err != nil {
		t.Errorf("cannot add segment: %v", err)
	}

	segment0Id, err := env.Repo.GetSegmentId(env.Ctx, segments[0])
	if err != nil {
		t.Errorf("cannot get segment id: %v", err)
	}

	if respAddSegment.Id != segment0Id {
		t.Errorf("cannot add segment got %d expected %d", respAddSegment.Id, segment0Id)
	}

	respAddDelete, err := env.Hdl.AddDeleteUserSegmentHandler(env.Ctx, handler.AddDeleteUserSegmentRequest{
		ToAdd: []handler.Segment{
			{
				Name:    segments[0],
				Expires: expires[0],
			},
			{
				Name:    segments[1],
				Expires: expires[1],
			},
			{
				Name:    segments[2],
				Expires: expires[2],
			},
		},
		ToDelete: nil,
		UserId:   userIds[0],
	})
	if err != nil {
		t.Errorf("cannot add delete user segmet pairs: %v", err)
	}

	if len(respAddDelete.AddedIds) != 3 {
		t.Errorf("not all user segment pairs were added: expected 3, got %d", len(respAddDelete.AddedIds))
	}

	respGetSegmentsOfUser, err := env.Hdl.GetSegmentsOfUserHandler(env.Ctx, handler.GetSegmentsOfUserRequest{Id: userIds[0]})
	if err != nil {
		t.Errorf("cannot get user segments: %v", err)
	}

	for _, segment := range respGetSegmentsOfUser.Segments {
		expires, err := time.Parse("2006-01-02T15:04:05", segment.Expires)
		if err != nil {
			t.Errorf("cannot parse date: %v", err)
		}
		if expires.Before(time.Now()) {
			t.Errorf("returned expires user segment pair %s expires %v", segment.Name, segment.Expires)
		}
	}

	_, err = env.Hdl.AddDeleteUserSegmentHandler(env.Ctx, handler.AddDeleteUserSegmentRequest{
		ToAdd:    nil,
		ToDelete: []string{segments[0], segments[1], segments[2]},
		UserId:   userIds[0],
	})
	if err != nil {
		t.Errorf("cannot add delete user segmet pairs: %v", err)
	}

	respGetSegmentsOfUser, err = env.Hdl.GetSegmentsOfUserHandler(env.Ctx, handler.GetSegmentsOfUserRequest{Id: userIds[0]})
	if err != nil {
		t.Errorf("cannot get user segments: %v", err)
	}
	if len(respGetSegmentsOfUser.Segments) != 0 {
		t.Errorf("wrong number of segments, expected 0, got %d", len(respGetSegmentsOfUser.Segments))
	}

	_, err = env.Hdl.DeleteSegmentHandler(env.Ctx, handler.DeleteSegmentRequest{Name: segments[0]})
	if err != nil {
		t.Errorf("cannot delete segment: %v", err)
	}
	_, err = env.Hdl.DeleteSegmentHandler(env.Ctx, handler.DeleteSegmentRequest{Name: segments[0]})
	if err == nil {
		t.Errorf("no error while deleting non-existing segment")
	}

	_, err = env.Hdl.FlushExpiredHandler(env.Ctx, struct{}{})
	if err != nil {
		t.Errorf("cannot flush expired user segment pairs: %v", err)
	}

}

const nameLength = 6

func genName() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	name := make([]byte, nameLength)
	for i := range name {
		name[i] = chars[rand.Intn(len(chars))]
	}

	return "test" + string(name)
}

func genDate() string {
	min := time.Now().AddDate(-1, 0, 0) // Subtract 1 year from current time
	max := time.Now().AddDate(1, 0, 0)  // Add 1 year to current time

	delta := max.Unix() - min.Unix()
	sec := rand.Int63n(delta)
	randomDate := min.Add(time.Duration(sec) * time.Second)
	return randomDate.Format("2006-01-02T15:04:05")
}
