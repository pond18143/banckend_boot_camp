package auth

import (
	"go-api-game-boot-camp/app"

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

func (ep *Endpoint) SignIn(c *gin.Context) {
	defer c.Request.Body.Close()

	var request credentials //model รับ input จาก body
	log.Info("[Input : ...]")

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	srv := NewAuthService(ep.CV, ep.EM)

	result, err := srv.EnCode(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if result.Value == "" {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Infof("Token : %+s", result.Value)

	c.JSON(http.StatusOK, result)
	return

}

func (ep *Endpoint) Register(c *gin.Context) {
	srv := NewRegisterService(ep.CV, ep.EM)
	defer c.Request.Body.Close()
	var request inputRegister
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("ShouldBindBodyWith : %s", err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Infoln("Register : ")
	log.Infoln("UserName : ", request.Username)
	log.Infoln("Email : ", request.Email)
	log.Infoln("SkinId : ", request.SkinId)
	log.Infoln("HatId : ", request.HatId)
	log.Infoln("ShirtId : ", request.ShirtId)
	log.Infoln("ShoesId : ", request.ShoesId)

	msg, err := srv.RegisterTransaction(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(msg.Status, msg)
	return
}
