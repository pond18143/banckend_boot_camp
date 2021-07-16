package loger

import (
	"go-api-game-boot-camp/app"
)

func createLog(request logModel) (err error) {

	if err = app.GameBootCamp.DB.
		Table("malar.dbo.log_application").
		Save(&request).Error; err != nil {
		return
	}

	return
}
