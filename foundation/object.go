package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "object.h"
import "C"
import (
	"sync"
	"unsafe"
)

// interface for all NSObject type
type Object interface {
	// return the delegate pointer of this object
	Ptr() unsafe.Pointer
	Dealloc()
}

var _ Object = (*NSObject)(nil)

// wrapper for NSObject
type NSObject struct {
	ptr unsafe.Pointer
}

func (o *NSObject) Dealloc() {

}

func (o *NSObject) Ptr() unsafe.Pointer {
	return o.ptr
}

func MakeObject(ptr unsafe.Pointer) *NSObject {
	if ptr == nil {
		return nil
	}
	return &NSObject{ptr}
}

var tasks = make(map[int64]func())
var taskLock sync.RWMutex
var currentTaskId int64

// AddDeallocHook add cocoa object dealloc hook
func AddDeallocHook(obj Object, hook func()) {
	if obj.Ptr() == nil {
		panic("cocoa pointer is nil")
	}
	taskLock.Lock()
	currentTaskId++
	id := currentTaskId
	tasks[id] = hook
	taskLock.Unlock()
	C.Dealloc_AddHook(obj.Ptr(), C.long(id))
}

//export runDeallocTask
func runDeallocTask(id int64) {

	taskLock.RLock()
	task := tasks[id]
	taskLock.RUnlock()

	task()

	taskLock.Lock()
	delete(tasks, id)
	taskLock.Unlock()
}
