package event

import (
	"github.com/hsiafan/cocoa/foundation/object"
)

type Event interface {
	object.Object
}

var _ Event = (*NSEvent)(nil)

type NSEvent struct {
	object.NSObject
}
