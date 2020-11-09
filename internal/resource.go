package internal

import (
	"github.com/hsiafan/cocoa/foundation/object"
	"sync"
	"sync/atomic"
)

// ResourceRegistry maintains resources, and their ids.
type ResourceRegistry struct {
	currentId int64
	resources map[int64]object.Object
	lock      sync.RWMutex
}

// NewResourceRegistry create new ResourceRegistry
func NewResourceRegistry() *ResourceRegistry {
	return &ResourceRegistry{
		resources: make(map[int64]object.Object),
	}
}

// NextId return next id for resource
func (r *ResourceRegistry) NextId() int64 {
	return atomic.AddInt64(&r.currentId, 1)
}

// Store store resource with it's id
func (r *ResourceRegistry) Store(id int64, obj object.Object) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.resources[id]; ok {
		panic("resource with this id already exists")
	}
	r.resources[id] = obj

}

// Store get resource by id
func (r *ResourceRegistry) Get(id int64) object.Object {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.resources[id]
}

// Store delete resource by id
func (r *ResourceRegistry) Delete(id int64) object.Object {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.resources[id]; ok {
		panic("resource with this id not exists")
	}
	obj := r.resources[id]
	delete(r.resources, id)
	return obj
}
