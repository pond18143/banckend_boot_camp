package lottery

import (
	"github.com/gin-gonic/gin/binding"
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
	srv := NewLotteryService(ep.CV, ep.EM)
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

//เรียกดู Lottery Filter ตามข้อมูลที่ส่งมา -> "status":1(ยังไม่ขาย),2(ขายแล้ว) "round_id":1 "lottery_number":"000"-"999" "character_id":11-16
//"paging_index":<1 "paging_size":<0
func (ep *Endpoint) LotteryStock(c *gin.Context) {
	defer c.Request.Body.Close()
	var request lotteryInput
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, err := srv.LotteryList(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, result)
		return
	}
	log.Info("Lottery stock")
	c.JSON(http.StatusOK, result)
	return
}


func (ep *Endpoint) RandomLottery(c *gin.Context) {
	defer c.Request.Body.Close()
	result,err := srv.RandLottery()
	if err != nil {
		//return err
		message := responseMessage{
			Status:             http.StatusBadRequest,
			MessageDescription: err.Error(), //record not found
		}
		log.Errorf("Error : %+v", message)
		c.JSON(http.StatusBadRequest, message)
		return
	}
	log.Info("Random Lottery")
	c.JSON(http.StatusOK, result)
	return
}
