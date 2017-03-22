// package log is a simple logger abstraction that buffers internally, and writes out at a user-defined period
// this enables us to log inside of an action, but delay where and when the log is written to an external source
// so that the http server and the "one and done" old school action invocations can use the same log, but present
// data back to the user in a simple manner.
package log

import (
	"bytes"
	"fmt"
	"io"
)

// Logger is a wrapper around a []string buffer. It is threadsafe questionmarks?
type Logger struct {
	buf bytes.Buffer
}

// NewLogger will return a new instance of Logger
func NewLogger() *Logger {
	return &Logger{}
}

// Info is a simple logging function which appends the log lines to the data
func (l *Logger) Info(line string) {
	l.buf.WriteString(line + `\n`)
}

// Infof is the format enabled version of Info
func (l *Logger) Infof(line string, vals ...interface{}) {
	l.buf.WriteString(fmt.Sprintf(line, vals...) + `\n`)
}

// Flush will flush the logger to the provided io.Writer. This could be stdout, stderr, a string builder, etc.
func (l *Logger) Flush(w io.Writer) (int64, error) {
	return l.buf.WriteTo(w)
}

// String will implement the Stringer interface and also return a copy of the log data as a string
func (l *Logger) String() string {
	return l.buf.String()
}
