package ping

import (
	"fmt"
	"net/http"

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
	log.Infof("Check Heartbeat : PingGetEndpoint")

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) PingGetParamsEndpoint(c *gin.Context) { //GET app/ping/:name
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : PingGetPramsEndpoint")

	//ดึงค่าจาก params ชื่อ name
	name := c.Params.ByName("name") //key name in param
	log.Infof("Params name : %s", name)

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result.Message = fmt.Sprintf("%s [%s]", result.Message, name)

	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) PingPostParamsAndBodyEndpoint(c *gin.Context) { //POST app/ping/:name
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : PingPostBodyEndpoint")

	//ดึงค่าจาก params ชื่อ name
	name := c.Params.ByName("name")
	log.Infof("Params name : %s", name)

	//ดึงค่าจาก body
	var request inputHeartbeat //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("Body Name : %s", request.Name)
	log.Infof("Body Age  : %d", request.Age)

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result.Message = fmt.Sprintf("%s ชื่อเล่น [%s] ชื่อเต็ม [%s] บริษัท [%d]", result.Message, name, request.Name, request.Age)

	//return success
	c.JSON(http.StatusOK, result)
	return
}
