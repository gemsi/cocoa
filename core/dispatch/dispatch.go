package dispatch

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
