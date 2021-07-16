package character

import (
	"go-api-game-boot-camp/service/auth"
	"net/http"

	"go-api-game-boot-camp/app"

	"github.com/gin-gonic/gin"
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
	srv := NewCharacterService(ep.CV, ep.EM)
	log.Infof("Check Heartbeat : characterGetEndpoint")

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


func (ep *Endpoint) GetCharacterInfo(c *gin.Context) { //GET app/ping
	defer c.Request.Body.Close()

	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()

	srv := NewCharacterService(ep.CV, ep.EM)
	log.Infof("GetCharacterInfo :")
	log.Infof("characterId : ", characterId)

	//เรียก logic
	result, err := srv.getCharacterInfo(characterId)
	log.Infof("result :%d",result)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}
