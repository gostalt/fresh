package logging

import (
	"fmt"
	"os"
)

type StdOut struct{}

func (l StdOut) Alert(p []byte) {
	l.log("Alert", p)
}

func (l StdOut) Critical(p []byte) {
	l.log("Critical", p)
}

func (l StdOut) Debug(p []byte) {
	l.log("Debug", p)
}

func (l StdOut) Emergency(p []byte) {
	l.log("Emergency", p)
}

func (l StdOut) Error(p []byte) {
	l.log("Error", p)
}

func (l StdOut) Info(p []byte) {
	l.log("Info", p)
}

func (l StdOut) Notice(p []byte) {
	l.log("Notice", p)
}

func (l StdOut) Warning(p []byte) {
	l.log("Warning", p)
}

func (l StdOut) log(level string, p []byte) {
	fmt.Fprintf(
		os.Stdout,
		"%s: %s\n",
		level,
		p,
	)
}
