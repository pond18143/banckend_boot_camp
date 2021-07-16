package farm

import (
	"net/http"

	"go-api-game-boot-camp/service/auth"

	"go-api-game-boot-camp/app"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
	EM *app.ErrorMessage
	CV *app.Configs
}

func NewEndpoint(conf *app.Configs, em *app.ErrorMessage) *Endpoint {
	return &Endpoint{
		CV: conf,
		EM: em,
	}
}

//รับ INPUT แปลงค่า

func (ep *Endpoint) PingGetEndpoint(c *gin.Context) { //GET app/ping
	defer c.Request.Body.Close()
	srv := NewFarmService(ep.CV, ep.EM)
	log.Infof("Check Heartbeat : farmGetEndpoint")

	//เรียก logic
	result, err := srv.checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) GetFarmInfo(c *gin.Context) { //GET app/ping
	defer c.Request.Body.Close()

	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()

	srv := NewFarmService(ep.CV, ep.EM)

	log.Infoln("GetFarmInfo :")
	log.Infoln("characterId : %v", characterId)

	//เรียก logic
	result, err := srv.getFarmInfo(characterId)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) Harvest(c *gin.Context) { //GET app/ping
	defer c.Request.Body.Close()

	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()

	srv := NewFarmService(ep.CV, ep.EM)

	log.Infof("Harvest :")

	var request harvestingRequest //model รับ input จาก body

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	checkpointX := request.CheckPointX
	checkpointY := request.CheckPointY
	log.Infoln("characterId : ", characterId)
	log.Infoln("checkpointX : ", checkpointX)
	log.Infoln("checkpointY : ", checkpointY)

	//เรียก logic
	result, err := srv.harvest(checkpointX, checkpointY, characterId)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) Planting(c *gin.Context) { //GET app/ping
	defer c.Request.Body.Close()
	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()

	srv := NewFarmService(ep.CV, ep.EM)
	log.Infof("Planting :")

	var request plantingRequest //model รับ input จาก body

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	itemId := request.ItemId
	checkpointX := request.CheckPointX
	checkpointY := request.CheckPointY
	log.Infof("characterId:%v", request.CharacterId)
	log.Infof("itemId:%v", request.ItemId)
	log.Infof("checkPointX:%v", request.CheckPointX)
	log.Infof("checlPointY:%v", request.CheckPointY)
	//เรียก logic

	result, err := srv.planting(characterId, itemId, checkpointX, checkpointY)
	log.Infof("result:%v", result)
	log.Infof("err:%v", err)

	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) Watering(c *gin.Context) { //GET app/farm
	defer c.Request.Body.Close()

	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()

	srv := NewFarmService(ep.CV, ep.EM)
	log.Infof(" Have you watered the plants yet ? :")

	//get input
	var request getFarmWatering

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	checkpointX := request.CheckPointX
	checkpointY := request.CheckPointY
	log.Infoln("characterId : ", characterId)
	log.Infoln("checkpointX : ", checkpointX)
	log.Infoln("checkpointY : ", checkpointY)

	//เรียก logic
	msg, err := srv.watering(characterId, checkpointX, checkpointY)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	} //return success
	c.JSON(http.StatusOK, msg)
	return
}
