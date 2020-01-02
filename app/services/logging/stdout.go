package logging

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
)

type StdOut struct{}

func (l StdOut) Alert(p []byte) {
	l.log("Alert", p, Red)
}

func (l StdOut) Critical(p []byte) {
	l.log("Critical", p, Red)
}

func (l StdOut) Debug(p []byte) {
	l.log("Debug", p, Yellow)
}

func (l StdOut) Emergency(p []byte) {
	l.log("Emergency", p, Yellow)
}

func (l StdOut) Error(p []byte) {
	l.log("Error", p, Red)
}

func (l StdOut) Info(p []byte) {
	l.log("Info", p, Green)
}

func (l StdOut) Notice(p []byte) {
	l.log("Notice", p, Bold)
}

func (l StdOut) Warning(p []byte) {
	l.log("Warning", p, Yellow)
}

func (l StdOut) log(level string, p []byte, color func(interface{}) Value) {
	fmt.Fprintf(
		os.Stdout,
		"%s %s\n",
		Bold(level+":"),
		color(p),
	)
}
