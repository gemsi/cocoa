package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "dispatch.h"
import "C"
import (
	"sync"
)

var tasks = make(map[int64]func())
var taskLock sync.RWMutex
var currentTaskId int64

// RunOnUIThread schedule task to run in ui thread, it does the same thing as
// 	dispatch_async(dispatch_get_main_queue(), ^{
// 	})
func RunOnUIThread(task func()) {
	taskLock.Lock()
	currentTaskId++
	id := currentTaskId
	tasks[id] = task
	taskLock.Unlock()
	C.Dispatch_RunOnUIThread(C.int(id))
}

//export runTask
func runTask(cid C.int) {
	id := int64(cid)

	taskLock.RLock()
	task := tasks[id]
	taskLock.RUnlock()

	task()

	taskLock.Lock()
	delete(tasks, id)
	taskLock.Unlock()
}
