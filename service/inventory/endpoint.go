package inventory

import (
	app "go-api-game-boot-camp/app"
	"go-api-game-boot-camp/service/auth"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
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

func (ep *Endpoint) GetInventory(c *gin.Context) {
	defer c.Request.Body.Close()

	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()
	srv := NewInventoryService(ep.CV, ep.EM)

	log.Infof("GetInventory : ")
	log.Infof("characterId : ", characterId)

	//call logic
	result, _, err := srv.getInventoryList(characterId) //characterId
	if err != nil {
		//return err
		log.Errorf("Error : %+v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//output of item list
	c.JSON(http.StatusOK, result)
	return
}
