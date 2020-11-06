package internal

import (
	"github.com/hsiafan/cocoa/foundation/notification"
	"sync"
)

// HandlerType for delegate Events
type HandlerType int

// HandlerRegistry make it easy to handle event
type HandlerRegistry struct {
	events map[HandlerType][]notification.Handler
	lock   sync.RWMutex
}

// NewHandlerRegistry create new HandlerRegistry
func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		events: make(map[HandlerType][]notification.Handler),
	}
}

// Add add new handler
func (r *HandlerRegistry) Add(event HandlerType, handler notification.Handler) {
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
func (r *HandlerRegistry) Get(event HandlerType) []notification.Handler {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.events[event]
}
