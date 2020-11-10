package widget

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Responder wrap cocoa NSResponder
type Responder interface {
	foundation.Object
	OnKeyDown(func(Event))
}

var _ Responder = (*NSResponder)(nil)

type NSResponder struct {
	foundation.NSObject
}

func MakeResponder(ptr unsafe.Pointer) *NSResponder {
	return &NSResponder{*foundation.MakeObject(ptr)}
}

func (*NSResponder) OnKeyDown(func(Event)) {

}
