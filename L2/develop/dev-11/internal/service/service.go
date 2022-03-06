package service

import (
	"dev-11/internal/domain/model"
	"dev-11/internal/repository/cache"
	"time"
)

type service struct {
	cache cache.Cache
}

func NewService(c cache.Cache) Service {
	return &service{cache: c}
}

func (s service) SetEvent(uuid int, event model.Event) {
	s.cache.Set(uuid, event)
}

func (s service) GetEvents(uuid int, t time.Time, duration time.Duration) []model.Event {
	res := make([]model.Event, 0)
	events := s.cache.Get(uuid)

	for _, event := range events {
		if event.Time.After(t) && event.Time.Before(t.Add(duration)) {
			res = append(res, event)
		}
	}
	return res
}

func (s service) DeleteEvent(uuid int, eventID int) {
	s.cache.Delete(uuid, eventID)
}

func (s service) IsValid(t string) (time.Time, error) {
	res, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}
