package handler

import (
	"time"

	"go-api-game-boot-camp/app"
	"go-api-game-boot-camp/service/loger"
	"go-api-game-boot-camp/service/sendgrid"

	"github.com/robfig/cron"
)

var log *loger.Loger

type Scheduler struct {
	EM *app.ErrorMessage
	CV *app.Configs
}

func InitScheduler(conf *app.Configs, em *app.ErrorMessage) *Scheduler{
	log = loger.NewLogController()
	return &Scheduler {
		EM: em,
		CV: conf,
	}
}

// enable_scheduler: true
// count down : "@every 25s"
// time on day : "TZ=Asia/Bangkok 59 23 * * *" #time_zone MM:HH
func (s *Scheduler)StartScheduler() {
	if s.CV.Scheduler.EnableJob {
		log.Info("Start Scheduler")
		s.startSendMailScheduler()
		// startCheckInBroadcastScheduler()
		// startAlertScheduler()
	}
}

func (s *Scheduler)startSendMailScheduler() {
	_timeScheduler := "@every 30s" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, s.sendMailJob)
	c.Start()
	log.Info("StartSendMailScheduler : ")
}

func (s *Scheduler)sendMailJob() {
	log.Info("##############################")
	log.Infof("Checking Pending Mail to Send : ", time.Now())
	log.Info("##############################")
	res, err := sendgrid.SendMail()
	if err != nil {
		log.Errorf("Error while sending mails", err)
	}
	log.Info(res)
}

// func startCheckInBroadcastScheduler() {
// 	_timeScheduler := "TZ=Asia/Bangkok 11 14 * * *" //time_zone MM:HH
// 	c := cron.New()
// 	c.AddFunc(_timeScheduler, broadcastJob)
// 	c.Start()
// 	msg := fmt.Sprintf("Schedule Boibot Send CheckIn Broadcast running at: %s", c.Entries()[0].Next.String())
// 	log.Infoln("StartCheckInBroadcastScheduler : ", msg)
// }

// func startAlertScheduler() {
// 	_timeScheduler := "@every 10s" //time_zone MM:HH
// 	c := cron.New()
// 	c.AddFunc(_timeScheduler, alertCheckInJob)
// 	c.Start()
// 	msg := fmt.Sprintf("Schedule Alert finish check in step running at: %s", c.Entries()[0].Next.String())
// 	log.Infoln("StartAlertScheduler : ", msg)
// }

// func broadcastJob() {
// 	log.Infoln("===============================")
// 	log.Infoln("Check In Broadcast Scheduler : ", time.Now())
// 	log.Infoln("===============================")

// }

// func alertCheckInJob() {
// 	log.Infoln("##############################")
// 	log.Infoln("Alert Check In Scheduler : ", time.Now())
// 	log.Infoln("##############################")

// }
