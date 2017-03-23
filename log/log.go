// Package log is a simple logger abstraction that buffers internally, and writes out at a user-defined period
// this enables us to log inside of an action, but delay where and when the log is written to an external source
// so that the http server and the "one and done" old school action invocations can use the same log, but present
// data back to the user in a simple manner.
package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

// Logger is a very lightweight interface around 2 simple log calls
type Logger interface {
	Println(string)
	Printf(string, ...interface{})
}

// BufferedLogger is a wrapper around a []string buffer. It is threadsafe questionmarks?
type BufferedLogger struct {
	buf bytes.Buffer
}

// NewBufferedLogger will return a new instance of Logger
func NewBufferedLogger() *BufferedLogger {
	return &BufferedLogger{}
}

// Println prints a line
func (l *BufferedLogger) Println(line string) {
	l.buf.WriteString(line + `\n`)
}

// Printf prints an f
func (l *BufferedLogger) Printf(line string, vals ...interface{}) {
	l.buf.WriteString(fmt.Sprintf(line, vals...))
}

// Flush will flush the logger to the provided io.Writer. This could be stdout, stderr, a string builder, etc.
func (l *BufferedLogger) Flush(w io.Writer) (int64, error) {
	return l.buf.WriteTo(w)
}

// String will implement the Stringer interface and also return a copy of the log data as a string
func (l *BufferedLogger) String() string {
	return l.buf.String()
}

// NormalLogger is just a passthrough to calling log.X
// The reason this exists instead of just using the builtin logger as an interface
// is that the Println signature takes a variadic of interfaces{} which i can't easily
// turn into strings for the buffer in the buffered logger. So, it's a little bit of
// fuffery to make things easier for the real purpose of the logger, which is buffering.
type NormalLogger struct{}

// NewNormalLogger returns a new normal logger
func NewNormalLogger() *NormalLogger {
	return &NormalLogger{}
}

// Println prints a line
func (l *NormalLogger) Println(line string) {
	log.Println(line)
}

// Printf prints an f
func (l *NormalLogger) Printf(line string, vals ...interface{}) {
	log.Printf(line, vals...)
}
