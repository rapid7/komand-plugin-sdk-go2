package dispatcher

import "github.com/komand/plugin-sdk-go2/message"

// NOOP will dispatch events to nowhere - it's for testing only
type NOOP struct{}

// NewNOOP returns a new NOOP dispatcher
func NewNOOP() *NOOP {
	return &NOOP{}
}

// Send dispatches a trigger event, except not for NOOP it doesn't
func (d *NOOP) Send(e *message.Response) error {
	return nil
}
