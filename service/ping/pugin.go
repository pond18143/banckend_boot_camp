package ping

import (
	"go-api-game-boot-camp/app"
)

func logHeartbeat(hb heartbeatModel) (err error) { //sql
	if err = app.GameBootCamp.DB.Table("document_header").Save(hb).Error; err != nil {
		return
	}
	return
}
