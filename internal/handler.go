package internal

import (
	"github.com/hsiafan/cocoa/foundation/notification"
	"sync"
)

// Event for delegate Events
type Event int

// HandlerRegistry make it easy to handle event
type HandlerRegistry struct {
	events map[Event][]notification.Handler
	lock   sync.RWMutex
}

// NewHandlerRegistry create new HandlerRegistry
func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		events: make(map[Event][]notification.Handler),
	}
}

// Add add new handler
func (r *HandlerRegistry) Add(event Event, handler notification.Handler) {
	if handler == nil {
		panic("handler is nil")
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	handlers := r.events[event]
	newHandlers := make([]notification.Handler, len(handlers)+1)
	copy(newHandlers, handlers)
	newHandlers[len(handlers)] = handler
	r.events[event] = newHandlers
}

// Get all handlers of the event
func (r *HandlerRegistry) Get(event Event) []notification.Handler {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.events[event]
}
