package logging

import (
	"fmt"
	"log"
	"os"
)

type File struct {
	file string
}

func MakeFile(path string) File {
	return File{
		file: path,
	}
}

func (l File) Alert(p []byte) {
	l.log("Alert", p)
}

func (l File) Critical(p []byte) {
	l.log("Critical", p)
}

func (l File) Debug(p []byte) {
	l.log("Debug", p)
}

func (l File) Emergency(p []byte) {
	l.log("Emergency", p)
}

func (l File) Error(p []byte) {
	l.log("Error", p)
}

func (l File) Info(p []byte) {
	l.log("Info", p)
}

func (l File) Notice(p []byte) {
	l.log("Notice", p)
}

func (l File) Warning(p []byte) {
	l.log("Warning", p)
}

func (l File) log(level string, p []byte) {
	f, err := os.OpenFile(l.file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(
		f,
		"%s: %s\n",
		level,
		p,
	)
}
