package event

import (
	"dev-11/internal/domain/model"
	"time"
)

type EventDTO struct {
	ID          int       `json:"ID"`
	UserID      int       `json:"user_ID"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Time        time.Time `json:"time"`
}

type DeleteEventDTO struct {
	ID     int `json:"ID"`
	UserID int `json:"user_ID"`
}

func (d EventDTO) toEvent() model.Event {
	return model.Event{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Time:        d.Time,
	}
}
