package auth

import (
	"go-api-game-boot-camp/app"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type registerRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *registerRepo) InitRegisterRepo(conf *app.Configs, em *app.ErrorMessage) *registerRepo {
	return &registerRepo{
		conf: conf,
		em:   em,
	}
}

type authRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *authRepo) InitAuthRepo(conf *app.Configs, em *app.ErrorMessage) *authRepo {
	return &authRepo{
		conf: conf,
		em:   em,
	}
}

func (repo *authRepo) getCharacter(request credentials) (detail character, err error) {

	if err = app.GameBootCamp.DB.Select("l.username ,l.password ,l.login_uuid ,c.id AS character_id ,c.login_id ,l.section").
		Table("bootcamp.dbo.character AS c ,bootcamp.dbo.login AS l").
		Where("l.username = ? AND l.id = c.login_id", request.Username).
		Find(&detail).Error; err != nil {
		return
	}

	return
}

func (repo *authRepo) updateLastLogin(request credentials, timeNow time.Time, section string, timeToken time.Time) (update roleLogin, err error) {
	if err = app.GameBootCamp.DB.Table("bootcamp.dbo.login").
		Where("username = ?", request.Username).
		Update(map[string]interface{}{
			"last_login":  timeNow,
			"section":     section,
			"exp_section": timeToken,
		}).
		Error; err != nil {
		return
	}
	update.LastLogin = timeNow
	update.ExpSection = timeToken
	return update, nil
}

func getSection(request int) (detail character, err error) {

	if err = app.GameBootCamp.DB.Select("l.username ,l.password ,l.login_uuid ,c.id AS character_id ,c.login_id ,l.section").
		Table("bootcamp.dbo.character AS c ,bootcamp.dbo.login AS l").
		Where("c.id = ? AND c.login_id = l.id ", request).
		Find(&detail).Error; err != nil {
		return
	}

	return
}

func (repo *registerRepo) checkEmail(request inputRegister) (result user, err error) {
	if err = app.GameBootCamp.DB.
		Table("bootcamp.dbo.login").
		Where("email = ?", request.Email).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) checkUsername(request inputRegister) (result user, err error) {
	if err = app.GameBootCamp.DB.
		Table("bootcamp.dbo.login").
		Where("username = ?", request.Username).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) createUser(tx *gorm.DB, request user) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.login").
		Create(&request).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) getUserId(tx *gorm.DB, request string) (user user, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.login").
		Where("username = ?", request).
		Find(&user).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) createCharacter(tx *gorm.DB, request inputCharacter) (character inputCharacter, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.character").
		Create(&request).Find(&character).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) getCharacterId(tx *gorm.DB, loginId int) (character inputCharacter, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.character").
		Where("login_id = ?", loginId).
		Find(&character).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) createInventory(tx *gorm.DB, request []inventory) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	for _, detail := range request {
		if err = tx.Table("bootcamp.dbo.inventory").
			Create(&detail).Error; err != nil {
			return
		}
	}
	return
}

func (repo *registerRepo) createFarm(tx *gorm.DB, request []farm) (msg MessageResponse, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	for _, detail := range request {
		if err = tx.Table("bootcamp.dbo.farm").
			Create(&detail).Error; err != nil {
			return
		}
	}
	msg.Status = http.StatusCreated
	msg.MessageCode = "0000"
	msg.MessageDescription = "Account created successfully."
	return
}
