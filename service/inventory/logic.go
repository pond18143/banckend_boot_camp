package inventory

import (
	app "go-api-game-boot-camp/app"

	log "github.com/sirupsen/logrus"
)

var srv *inventoryService

type inventoryService struct {
	conf *app.Configs
	em   *app.ErrorMessage
	repo *inventoryRepo
}

func Init(conf *app.Configs, em *app.ErrorMessage) {
	srv = &inventoryService{
		conf: conf,
		em:   em,
		repo: &inventoryRepo{conf: conf},
	}
}

func NewInventoryService(conf *app.Configs, em *app.ErrorMessage) *inventoryService {
	repo := inventoryRepo{}

	return &inventoryService{
		conf: conf,
		em:   em,
		repo: repo.InitInventoryRepo(conf, em),
	}
}

func (srv *inventoryService) getInventoryList(request int) (result resInventory, message responseMessage, err error) {
	//call header pugin
	var header headerInventory
	header, err = inventoryHeader(request)
	if err != nil {
		err = srv.em.Character.CharacterIDNotFound
		log.Errorln("srv.repo.getInventoryList fail : ", err.Error())
		return
	}
	log.Infof("document_header : %+v", header)

	//call detail pugin
	var detail []DetailInventory
	detail, err = inventoryDetail(request)
	//check err
	if err != nil {
		err = srv.em.Character.GetItemListNotFound
		log.Errorln("srv.repo.getInventoryList fail : ", err.Error())
		return
	}
	log.Infof("document_detail : %+v", detail)

	//if success show character id and item list
	result = resInventory{
		header,
		detail,
	}
	return
}
