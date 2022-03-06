package cache

import "dev-11/internal/domain/model"

type Cache interface {
	Set(uuid int, event model.Event)
	Get(uuid int) []model.Event
	Delete(uuid int, eventID int)
}
