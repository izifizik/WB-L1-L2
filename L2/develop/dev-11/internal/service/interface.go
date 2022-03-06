package service

import (
	"dev-11/internal/domain/model"
	"time"
)

type Service interface {
	SetEvent(uuid int, event model.Event)
	GetEvents(uuid int, time time.Time, duration time.Duration) []model.Event
	DeleteEvent(uuid int, eventID int)

	IsValid(t string) (time.Time, error)
}
