package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "base.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"sync"
)

var tasks = make(map[int64]func())
var taskLock sync.RWMutex
var currentTaskId int64

// MainQueueAsync schedule task to run in main thread async, it does the same thing as
// 	dispatch_async(dispatch_get_main_queue(), ^{
// 	})
// This method can be used to do ui operations in UI thread, for appkit etc..
func MainQueueAsync(task func()) {
	taskLock.Lock()
	currentTaskId++
	id := currentTaskId
	tasks[id] = task
	taskLock.Unlock()
	C.Dispatch_MainQueueAsync(C.long(id))
}

// WithAutoreleasePool run code in a new auto release pool.
func WithAutoreleasePool(task func()) {
	taskLock.Lock()
	currentTaskId++
	id := currentTaskId
	tasks[id] = task
	taskLock.Unlock()
	C.Run_WithAutoreleasePool(C.long(id))
}

// AddDeallocHook add cocoa object dealloc hook
func AddDeallocHook(obj foundation.Object, hook func()) {
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

//export runTask
func runTask(id int64) {

	taskLock.RLock()
	task := tasks[id]
	taskLock.RUnlock()

	task()

	taskLock.Lock()
	delete(tasks, id)
	taskLock.Unlock()
}
