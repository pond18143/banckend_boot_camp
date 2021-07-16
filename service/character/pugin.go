package character

import (
	"go-api-game-boot-camp/app"
)

type characterRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *characterRepo) InitCharacterRepo(conf *app.Configs, em *app.ErrorMessage) *characterRepo {
	return &characterRepo{
		conf: conf,
		em:   em,
	}
}

func (repo *characterRepo) logHeartbeat(hb heartbeatModel) (err error) { //sql
	if err = app.GameBootCamp.DB.Table("document_header").Save(hb).Error; err != nil {
		return
	}
	return
}

func (repo *characterRepo) getCharacterInfoByCharacterId(characterId int) (result characterDetail, err error) {

		if err = app.GameBootCamp.DB.Select("cc.id,ll.username,cc.login_id,cc.gold,cc.coin,cc.gender,cc.skin_id,cc.hat_id,cc.shirt_id,cc.shoes_id,cc.update_date").
			Table("bootcamp.dbo.character AS cc").
			Joins("INNER JOIN bootcamp.dbo.login AS ll ON ll.id = cc.login_id").
			Where("cc.id = ?", characterId).
			Find(&result).Error; err != nil {
			return
			}
	return
}
func (repo *characterRepo) getBuffInfoByCharacterId(characterId int) (result []buffDetail, err error) {
	if err = app.GameBootCamp.DB.Table("buff").Select("character_id,buff_name,remaining,start_date,end_date,update_date").Where("character_id=?", characterId).Find(&result).Error; err != nil {
		return
	}
	return
}
