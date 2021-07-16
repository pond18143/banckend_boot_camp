package auth

import (
	"time"
	"strings"
	"go-api-game-boot-camp/app"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

var rs *registerService

type registerService struct {
	conf *app.Configs
	em   *app.ErrorMessage
	repo *registerRepo
}

func NewRegisterService(conf *app.Configs, em *app.ErrorMessage) *registerService {
	repo := registerRepo{}

	return &registerService{
		conf: conf,
		em:   em,
		repo: repo.InitRegisterRepo(conf, em),
	}
}

func (rs *registerService) RegisterTransaction(request inputRegister) (msg MessageResponse, err error) {
	log.Info("[Register Validation]")
	request.Username = strings.TrimSpace(request.Username)

	//validate username empty
	if(request.Username == "") {
		err = rs.em.Register.ValidateFail.ValidateUsernameEmpty
		log.Error("RegisterTransaction Failed : " + err.Error())
		return
	}

	//validate username max 15
	if len(request.Username) > 15 {
		err = rs.em.Register.ValidateFail.ValidateUsernameLength
		log.Error("RegisterTransaction Failed : " + err.Error())
		return
	}
	//validate password max 20
	if len(request.Password) > 20 || len(request.RepeatedPassword) > 20 {
		err = rs.em.Register.ValidateFail.ValidatePasswordLength
		log.Error("RegisterTransaction Failed : " + err.Error())
		return
	}
	//validate email max 30
	if len(request.Email) > 30 {
		err = rs.em.Register.ValidateFail.ValidateEmailLength
		log.Error("RegisterTransaction Failed : " + err.Error())
		return
	}

	//validate match password
	if request.Password != request.RepeatedPassword {
		err = rs.em.Register.ValidateFail.PasswordNotMatch
		log.Error("RegisterTransaction Failed : " + err.Error())
		return
	}

	//validate duplicate email
	var userFound user
	userFound, err = rs.repo.checkEmail(request)
	if (userFound != user{}) || (err != nil && err != gorm.ErrRecordNotFound) {
		err = rs.em.Register.ValidateFail.DuplicateEmail
		log.Error("RegisterTransaction Failed : " + err.Error())
		return
	}

	//validate duplicate username
	userFound, err = rs.repo.checkUsername(request)
	if (userFound != user{}) || (err != nil && err != gorm.ErrRecordNotFound) {
		err = rs.em.Register.ValidateFail.DuplicateUsername
		log.Error("RegisterTransaction Failed : " + err.Error())
		return
	}

	//create transaction -> user, character, inventory, farm
	if err == gorm.ErrRecordNotFound {
		log.Info("[-----------------Transaction Started-----------------]")
		tx := app.GameBootCamp.DB.Begin()
		err = nil
		loginUuid := uuid.New().String()
		password := encryptPassword(compilePassword(request.Password, loginUuid))
		var user = user{
			LoginUuid: loginUuid,
			Username:  request.Username,
			Password:  password,
			Email:     request.Email,
		}
		time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		user.RegisterDate = time
		user.CreateDate = time
		user.UpdateDate = time

		//create user
		log.Info("[Create User]")
		err = rs.repo.createUser(tx, user)
		if err != nil {
			err = rs.em.Register.CreateFail.CreateUsername
			log.Error("RegisterTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}

		//get user_id for character_id
		user, err = rs.repo.getUserId(tx, string(request.Username))
		if err != nil {
			err = rs.em.Register.CreateFail.CreateUsername
			log.Error("RegisterTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}

		var character = inputCharacter{
			LoginId: user.Id,
			Gender:  request.Gender,
			SkinId:  request.SkinId,
			HatId:   request.HatId,
			ShirtId: request.ShirtId,
			ShoesId: request.ShoesId,
		}
		character.CreateDate = time
		character.UpdateDate = time

		//create character
		log.Info("[Create Character]")
		character, err = rs.repo.createCharacter(tx, character)
		if err != nil {
			err = rs.em.Register.CreateFail.CreateCharacter
			log.Error("RegisterTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}

		//get character id for another table
		character, err = rs.repo.getCharacterId(tx, user.Id)
		if err != nil {
			err = rs.em.Register.CreateFail.CreateCharacter
			log.Error("RegisterTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}

		var inventory = []inventory{
			{
				CharacterId:	character.Id,	
				ItemId:			6,
				Quantity:		3,	
			},
			{
				CharacterId:	character.Id,	
				ItemId:			29,
				Quantity:		999999,	
			},					
		}
		inventory[0].CreateDate = time
		inventory[0].UpdateDate = time
		inventory[1].CreateDate = time
		inventory[1].UpdateDate = time

		//create inventory
		log.Info("[Create Inventory]")
		err = rs.repo.createInventory(tx, inventory)
		if err != nil {
			err = rs.em.Register.CreateFail.CreateInventory
			log.Error("RegisterTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}

		var farm = []farm {
			{ CharacterId:	character.Id, CheckPointX:	1, CheckPointY: 1, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	1, CheckPointY: 2, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	1, CheckPointY: 3, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	2, CheckPointY: 1, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	2, CheckPointY: 2, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	2, CheckPointY: 3, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	3, CheckPointY: 1, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	3, CheckPointY: 2, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	3, CheckPointY: 3, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	4, CheckPointY: 1, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	4, CheckPointY: 2, CreateDate: time, UpdateDate: time },
			{ CharacterId:	character.Id, CheckPointX:	4, CheckPointY: 3, CreateDate: time, UpdateDate: time },
		}

		//create farm
		log.Info("[Create Farm]")
		msg, err = rs.repo.createFarm(tx, farm)
		if err != nil {
			err = rs.em.Register.CreateFail.CreateFarm
			log.Error("RegisterTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}
		tx.Commit()
	}
	return
}
