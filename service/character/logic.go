package character

import (
	log "github.com/sirupsen/logrus"
	"go-api-game-boot-camp/app"
	"time"
)

var srv *characterService

type characterService struct {
	conf *app.Configs
	em   *app.ErrorMessage
	repo *characterRepo
}

func Init(conf *app.Configs, em *app.ErrorMessage) {
	srv = &characterService{
		conf: conf,
		em:   em,
		repo: &characterRepo{conf: conf},
	}
}

func NewCharacterService(conf *app.Configs, em *app.ErrorMessage) *characterService {
	repo := characterRepo{}

	return &characterService{
		conf: conf,
		em:   em,
		repo: repo.InitCharacterRepo(conf, em),
	}
}

func (srv *characterService) checkHeartbeat() (result heartbeatModel, err error) {

	result.Message = "character"
	result.DateTime = time.Now()

	//err = logHeartbeat(result)
	//if err != nil {
	//	return
	//}

	return
}

func (srv *characterService) getCharacterInfo(characterId int) (result characterDetailRes, err error) {
	log.Infof("getCharacterInfo :")
	//get Farm info by character_id
	var characterInfo characterDetail
	characterInfo, err = srv.repo.getCharacterInfoByCharacterId(characterId)
	if err != nil {
		log.Errorln("srv.repo.getCharacterInfoByCharacterId fail : ", err)
		err = srv.em.Farm.GetCharacterInfoNotFound
		log.Errorln("srv.repo.getCharacterInfoByCharacterId fail : ", err)
		return
	}
	log.Infof("characterInfo :%d", characterInfo)
	var buffInfo []buffDetail
	buffInfo, err = srv.repo.getBuffInfoByCharacterId(characterId)
	if err != nil {
		log.Errorln("srv.repo.getBuffInfoByCharacterId fail : ", err)
		err = srv.em.Farm.GetBuffNotFound
		log.Errorln("srv.repo.getBuffInfoByCharacterId fail : ", err)
		return
	}
	log.Infof("buffInfo :%d", buffInfo)

	var buffData []buffDetail

	for _, value := range buffInfo {
		t := time.Now()

		if inTimeSpan(value.StartDate, value.EndDate, t) {

			buffState := buffDetail{

				BuffName:    value.BuffName,
				Remaining:   value.Remaining,
				StartDate:   value.StartDate,
				EndDate:     value.EndDate,
				UpdateDate:  value.UpdateDate,
				Description: value.Description,
				Value:       value.Value,
			}
			buffData = append(buffData, buffState)
		}
	}

	result = characterDetailRes{
		characterInfo,
		buffData,
	}
	return
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
