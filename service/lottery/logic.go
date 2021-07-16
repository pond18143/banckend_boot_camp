package lottery

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"

	"go-api-game-boot-camp/app"
)

var srv *lotteryService

type lotteryService struct {
	conf *app.Configs
	em   *app.ErrorMessage
	repo *lotteryRepo
}

func Init(conf *app.Configs, em *app.ErrorMessage) {
	srv = &lotteryService{
		conf: conf,
		em:   em,
		repo: &lotteryRepo{conf: conf},
	}
}

func NewLotteryService(conf *app.Configs, em *app.ErrorMessage) *lotteryService {
	repo := lotteryRepo{}

	return &lotteryService{
		conf: conf,
		em:   em,
		repo: repo.InitLotteryRepo(conf, em),
	}
}

func (srv *lotteryService) checkHeartbeat() (result heartbeatModel, err error) {

	result.Message = "lottery"
	result.DateTime = time.Now()

	//err = logHeartbeat(result)
	//if err != nil {
	//	return
	//}

	return
}

func (srv *lotteryService)LotteryList(request lotteryInput) (result lotteryOutput, err error) {
	log.Infof("result: %+v", request)
	if request.RoundId == 0 {
		request.RoundId = 1 //หากไม่รู้ RoundId ให้ใช้เป็น 1
		log.Infof("result: %+v", request)
	}
	result, err = LotteryAll(request)
	if err != nil {
		err=srv.em.Lottery.GetLotteryError
		log.Errorln("srv.repo.LotteryAll fail : ", err)
		return
	}
	log.Infof("result: %+v", result)
	return
}

func (srv *lotteryService)RandLottery() (result exchangGamu,err error) {
	count,err := LotteryCount()
	if err != nil {
		err=srv.em.Lottery.GetLotteryCountError
		log.Errorln("srv.repo.LotteryCount fail : ", err)
		return
	}
	if count.SumQuantity == 0 {
		result.Status="successful"
		result.Description="exchange 0 item"
		return
	}
	log.Infof("result: %+v", result)
	CharId,err := CheckGamu()
	if err != nil {
		err=srv.em.Lottery.GetGamuOwnerError
		log.Errorln("srv.repo.CheckGamu fail : ", err)
		return
	}
	log.Infof("result: %+v", CharId)
	LotNum,err := RandomLottery(count.SumQuantity)
	if err != nil {
		err=srv.em.Lottery.GetRandomLotteryError
		log.Errorln("srv.repo.RandomLottery fail : ", err)
		return
	}
	log.Infof("result: %+v", LotNum)
	err=AddCharId(CharId,LotNum,count)
	if err != nil {
		err=srv.em.Lottery.UpdateLotteryOwnerError
		log.Errorln("srv.repo.AddCharId fail : ", err)
		return
	}
	result.Status="successful"
	result.Description="exchange "+strconv.Itoa(count.SumQuantity)+" item"
	return
}
