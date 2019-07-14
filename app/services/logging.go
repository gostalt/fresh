package services

import (
	"gostalt/config"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/sarulabs/di"
	jww "github.com/spf13/jwalterweatherman"
)

type LoggingServiceProvider struct {
	BaseServiceProvider
}

func (p LoggingServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "logger",
		Build: func(c di.Container) (interface{}, error) {
			log.Println("getting log writer")
			logWriter := p.getLogWriter()

			return jww.NewNotepad(
				jww.LevelInfo,
				jww.LevelTrace,
				ioutil.Discard,
				logWriter,
				config.Get("app", "name"),
				log.Ldate|log.Ltime,
			), nil
		},
	})
}

func (p LoggingServiceProvider) getLogWriter() (f io.Writer) {
	logType := config.Get("logging", "driver")

	switch logType {
	case "single":
		f = p.singleLog()
	case "daily":
		f = p.dailyLog()
	case "stdout":
		f = os.Stdout
	default:
		f = ioutil.Discard
	}

	return
}

// singleLog returns a file that all logs are saved to. The file
// is not recreated or truncated—it is added to with each log.
func (p LoggingServiceProvider) singleLog() io.Writer {
	return p.fileLog("gostalt")
}

// TODO: This doesn't really work at the moment - it creates a
// log file for the date in question when the app is created,
// as a one time operation: needs to be checked at log time.
func (p LoggingServiceProvider) dailyLog() io.Writer {
	date := time.Now().Format("2006-02-01")
	return p.fileLog(date + "-gostalt")
}

func (p LoggingServiceProvider) fileLog(name string) (f io.Writer) {
	logDir := config.Get("logging", "dir")

	f, err := os.OpenFile(
		logDir+name+".log",
		os.O_RDWR|os.O_APPEND|os.O_CREATE,
		0666,
	)
	if err != nil {
		panic("unable to use `single` log driver")
	}

	return
}
