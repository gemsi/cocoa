package widget

import (
	"github.com/hsiafan/cocoa/foundation"
)

type Event interface {
	foundation.Object
}

var _ Event = (*NSEvent)(nil)

type NSEvent struct {
	foundation.NSObject
}
