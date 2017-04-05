// Package cli contains some basic functions, "constants", and helpers for the main
// driver of the plugin
package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	plog "github.com/komand/plugin-sdk-go2/log"
	"github.com/komand/plugin-sdk-go2/message"
	ansi "github.com/mgutz/ansi"
)

// These define various color constants
var (
	Lime  = ansi.ColorCode("green+h:black")
	Red   = ansi.ColorCode("red")
	Green = ansi.ColorCode("green")
	Reset = ansi.ColorCode("reset")
)

// Args represents a simplified set of data passed in from the cli
type Args struct {
	Command     string
	SubCommands []string
	Port        int
}

// GetArgsFromCLI specially picks apart os.Args for the specific ways the engine or the cli
// might invoke a plugin.
func GetArgsFromCLI() (*Args, error) {
	fmt.Println(os.Args)
	if len(os.Args) <= 1 {
		return nil, errors.New("you must specify command to invoke a plugin")
	}
	// First, we have one optional arg - look for it
	args := &Args{}
	for i, arg := range os.Args {
		// If this is the arg for port, and there is another arg after it
		if arg == "--port" {
			// Alsp, remove these from os.Args afterwards, so that the rest of the pos-params are fine
			if len(os.Args) > i {
				var err error
				args.Port, err = strconv.Atoi(os.Args[i])
				if err != nil {
					return nil, fmt.Errorf("you provided an invalid value for --port: %s", os.Args[i])
				}
				os.Args = append(os.Args[:i], os.Args[i+2:]...)
			} else {
				os.Args = append(os.Args[:i], os.Args[i:]...)
			}
			break
		}
	}
	args.Command = os.Args[1]
	args.SubCommands = os.Args[2:]
	return args, nil
}

// HandleShutdown will block on the os.Signal channel, then begin a shutdown procedure
// by signalling the other goroutines via the cancellation context.
func HandleShutdown(cancel context.CancelFunc, c chan os.Signal) {
	<-c
	cancel()
	// Major hack - we should be blocking on a channel or series of channels that say
	// this part of the runtime has shut down, it's ok to proceed, then exit
	// TODO determine if we need 1 or many channels to wait on, pass (it/them) in, and wait
	// before os.Exitting
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

// WrapTriggerTestResult does the boring 1-off work of taking a trigger test output
// and wrapping it in the message envelopes and setting statuses. We only need this
// for triggers tests, since regular triggers submit data with the wrapper in another
// path. Additionally, actions also already wrap their outputs for their own purposes.
// TODO see if there is a simple way to combine all of those to just use this?
func WrapTriggerTestResult(log plog.Logger, o interface{}, err error) *message.V1 {
	l := ""
	if blog, ok := log.(*plog.BufferedLogger); ok {
		l = blog.String()
	}
	response := &message.Response{
		Meta:   []byte(`{}`),
		Output: o,
		Status: "ok",
		Log:    l,
	}
	if err != nil {
		response.Error = err.Error()
		response.Status = "error"
	}
	wrapper := &message.V1{
		Body:    response,
		Type:    "trigger_event",
		Version: "v1",
	}
	return wrapper
}
