package jobs

import (
	"github.com/gostalt/logger"
)

// Quote displays a message to the user each time they run the
// `./gostalt schedule` command. It can be deleted.
type Quote struct {
	Logger logger.Logger
}

func (j Quote) message() []byte {
	return []byte("I'm a scheduled job that comes bundled with Gostalt. You can find me in app/jobs/Quote.go")
}

func (j Quote) ShouldFire() bool {
	// Returning true here means that this job will run every
	// time the schedule is activated.
	return true
}

func (j Quote) Handle() error {
	j.Logger.Info(j.message())
	return nil
}
