package responder

import (
	"github.com/hsiafan/cocoa/foundation/object"
	"unsafe"

	"github.com/hsiafan/cocoa/appkit/event"
)

// Responder wrap cocoa NSResponder
type Responder interface {
	object.Object
	OnKeyDown(func(event.Event))
}

var _ Responder = (*NSResponder)(nil)

type NSResponder struct {
	object.NSObject
}

func Make(ptr unsafe.Pointer) *NSResponder {
	return &NSResponder{*object.Make(ptr)}
}

func (*NSResponder) OnKeyDown(func(event.Event)) {

}
