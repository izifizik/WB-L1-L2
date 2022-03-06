package model

import "time"

type Event struct {
	ID          int
	Title       string
	Description string
	Time        time.Time
}
