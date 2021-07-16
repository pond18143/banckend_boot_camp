package market

import (
	"go-api-game-boot-camp/app"
	"go-api-game-boot-camp/service/auth"
	"net/http"

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
	srv := NewMarketService(ep.CV, ep.EM)
	log.Infof("Check Heartbeat : market")

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

func (ep *Endpoint) SellItem(c *gin.Context) {
	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()
	
	srv := NewMarketService(ep.CV, ep.EM)
	defer c.Request.Body.Close()

	var request inputSellItem
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	itemId := request.ItemId
	quantity := request.Quantity

	msg, err := srv.SellItemTransaction(characterId, itemId, quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(msg.Status, msg)
	return
}

func (ep *Endpoint) BuyItem(c *gin.Context) {
	defer c.Request.Body.Close()
	srv := NewMarketService(ep.CV, ep.EM)
	log.Infof("[Buy Item : market]")

	//model รับ input จาก body
	var request BuyItemRequest
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// get characterId
	authService := auth.NewAuthService(ep.CV, ep.EM)
	authService.GetClaimCurrent(c)
	characterId := authService.GetCharacterIdInClaim()

	//เรียก logic
	err := srv.BuyItem(characterId, request)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//return success
	c.JSON(http.StatusOK, messageResponse{
		Status:             http.StatusCreated,
		MessageCode:        "0000",
		MessageDescription: "Buy item successfully.",
	})
	return
}
