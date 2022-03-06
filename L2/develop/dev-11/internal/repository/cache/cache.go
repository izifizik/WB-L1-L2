package cache

import (
	"dev-11/internal/domain/model"
	"sync"
)

type cache struct {
	cache map[int][]model.Event
}

var m sync.Mutex

//NewCache - create new Cache interface reference
func NewCache() Cache {
	cacheMap := make(map[int][]model.Event)
	return &cache{cache: cacheMap}
}

//Set - set value to cache (map)
func (c *cache) Set(uuid int, event model.Event) {
	m.Lock()
	events := c.cache[uuid]
	for i, e := range events {
		if e.ID == event.ID {
			e = event
			events[i] = e
			m.Unlock()
			return
		}
	}
	c.cache[uuid] = append(c.cache[uuid], event)
	m.Unlock()
}

//Get - get value from cache (map)
func (c cache) Get(uuid int) []model.Event {
	return c.cache[uuid]
}

func (c *cache) Delete(uuid int, eventID int) {
	m.Lock()
	events := c.cache[uuid]
	for i, event := range events {
		if event.ID == eventID {
			events = remove(events, i)
			break
		}
	}
	c.cache[uuid] = events
	m.Unlock()
}

func remove(slice []model.Event, s int) []model.Event {
	return append(slice[:s], slice[s+1:]...)
}
