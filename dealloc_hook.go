package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "dealloc_hook.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"sync"
)

var deallocHooks = make(map[int64]func())
var deallocHookLock sync.RWMutex
var deallocHookId int64

// AddDeallocHook add cocoa object dealloc hook
func AddDeallocHook(obj foundation.Object, hook func()) {
	deallocHookLock.Lock()
	deallocHookId++
	id := deallocHookId
	deallocHooks[id] = hook
	deallocHookLock.Unlock()
	C.Dealloc_AddHook(obj.Ptr(), C.int(id))
}

//export runDeallocHook
func runDeallocHook(hookId int64) {
	id := int64(hookId)

	deallocHookLock.RLock()
	task := deallocHooks[id]
	deallocHookLock.RUnlock()

	task()

	deallocHookLock.Lock()
	delete(deallocHooks, id)
	deallocHookLock.Unlock()
}
