package loger

import (
	//"strings"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)


type Loger struct{
	id string
}


func NewLogController() *Loger {

		logControllerPointer := &Loger{
			id: uuid.New().String(),
		}

	return logControllerPointer

}

func (l Loger) Info (message string) (err error) {
	log.Info(message)
	var log = logModel{
		SeriesId : l.id,
		LogLevel : InfoLevel,
		Message : message,
		DateTime : time.Now(),
	}


	err = createLog(log)
	if err != nil {

		return
	}

	return
}

func (l Loger) Debug (message string) (err error) {
	log.Debug(message)
	var log = logModel{
		SeriesId : l.id,
		LogLevel : DebugLevel,
		Message : message,
		DateTime : time.Now(),

	}
	err = createLog(log)
	if err != nil {

		return
	}

	return
}

func (l Loger) Error(message string) (err error) {
	log.Error(message)
	var log = logModel{
		SeriesId : l.id,
		LogLevel : ErrorLevel,
		Message : message,
		DateTime : time.Now(),

	}
	err = createLog(log)
	if err != nil {

		return
	}

	return
}

func (l Loger) Infof (format string, args interface{}) (err error) {
	log.Infof(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId : l.id,
		LogLevel : InfoLevel,
		Message : message,
		DateTime : time.Now(),
	}
	err = createLog(log)
	if err != nil {

		return
	}

	return
}

func (l Loger) Debugf (format string, args interface{}) (err error) {
	log.Debugf(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId : l.id,
		LogLevel : DebugLevel,
		Message : message,
		DateTime : time.Now(),

	}
	err = createLog(log)
	if err != nil {

		return
	}

	return
}

func (l Loger) Errorf(format string, args interface{}) (err error) {
	log.Errorf(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId : l.id,
		LogLevel : ErrorLevel,
		Message : message,
		DateTime : time.Now(),

	}
	err = createLog(log)
	if err != nil {

		return
	}

	return
}
