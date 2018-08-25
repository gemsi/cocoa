package responder

import (
	"unsafe"

	"github.com/hsiafan/cocoa/event"
	"github.com/hsiafan/cocoa/foundation"
)

// Responder wrap cocoa NSResponder
type Responder interface {
	foundation.Object
	OnKeyDown(func(event.Event))
}

var _ Responder = (*NSResponder)(nil)

type NSResponder struct {
	foundation.NSObject
}

func Make(ptr unsafe.Pointer) *NSResponder {
	return &NSResponder{*foundation.MakeObject(ptr)}
}

func (*NSResponder) OnKeyDown(func(event.Event)) {

}
