package farm

import (
	"net/http"
	"time"

	"go-api-game-boot-camp/app"

	"github.com/jinzhu/gorm"
)

type farmRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *farmRepo) InitFarmRepo(conf *app.Configs, em *app.ErrorMessage) *farmRepo {
	return &farmRepo{
		conf: conf,
		em:   em,
	}
}

func (repo *farmRepo) getFarmInfoListByCharacterId(characterId int) (result []farm, err error) {
	if err = app.GameBootCamp.DB.Table("farm").Select("*").Where("character_id = ?", characterId).Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) getFarmInfoByXY(checkPointX, checkPointY, characterId int) (result farm, err error) {
	if err = app.GameBootCamp.DB.Table("farm").Select("*").Where("check_point_x = ? AND check_point_y = ? AND character_id = ?", checkPointX, checkPointY, characterId).Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) updateHarvestedInFarm(tx *gorm.DB, checkPointX, checkPointY, remainingHarvest, characterId int) (err error) {

	if tx == nil {
		tx = app.GameBootCamp.DB
	}

	if remainingHarvest == 0 { //reset
		if err = tx.Table("farm").Where("check_point_x = ? AND check_point_y = ? AND character_id = ?", checkPointX, checkPointY, characterId).Update(map[string]interface{}{
			"remaining_Harvest": 0,
			"plant_dex_id":      0,
			"is_watered":        false,
			"plant_date":        time.Time{}.Truncate(24*time.Hour),
			"harvest_date":      time.Time{},
			"update_date":       time.Now(),
		}).Error; err != nil {
			return
		}
	} else {
		if err = tx.Table("farm").Where("check_point_x = ? AND check_point_y = ? AND character_id = ?", checkPointX, checkPointY, characterId).Update(map[string]interface{}{
			"remaining_Harvest": remainingHarvest,
			"plant_date":       time.Now(),
			"update_date":       time.Now(),
		}).Error; err != nil {
			return
		}
	}

	return
}

func (repo *farmRepo) getInventoryByCharacterIdAndItemId(itemId, characterId int) (result inventory, err error) {
	if err = app.GameBootCamp.DB.Table("inventory").Select("*").Where("item_id = ? AND character_id = ?", itemId, characterId).Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) updateInventoryByCharacterId(tx *gorm.DB, itemId, quantity, characterId int) (err error) {

	if tx == nil {
		tx = app.GameBootCamp.DB
	}

	if err = tx.Table("inventory").Where("item_id = ? AND character_id = ?", itemId, characterId).Update(map[string]interface{}{
		"quantity":    quantity,
		"update_date": time.Now(),
	}).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) addNewItemToInventoryByCharacterId(tx *gorm.DB, inventory inventory) (err error) {

	if tx == nil {
		tx = app.GameBootCamp.DB
	}

	if err = tx.Table("inventory").Create(inventory).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) getPlantDexById(id int) (result plantDex, err error) {
	if err = app.GameBootCamp.DB.Table("plant_dex").Select("*").Where("id = ?", id).Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) getPlantDexListByPlantType(fruitType FruitType, stateName *string, plantName *string, isSpecialItem *bool) (result []plantDex, err error) {
	tx := app.GameBootCamp.DB.Table("plant_dex").Select("*").Where("plant_type = ?", int(fruitType))

	if stateName != nil {
		tx = tx.Where("state_name = ?", *stateName)
	}

	if plantName != nil {
		tx = tx.Where("plant_name = ?", *plantName)
	}

	if isSpecialItem != nil {
		tx = tx.Where("is_special_item = ?", *isSpecialItem)
	}

	if err = tx.Order("hour_to_grow desc").Find(&result).Error; err != nil {
		return
	}

	return
}

func (repo *farmRepo) getPlantDexListByStateName(fruitType FruitType, isSpecialItem bool) (result []plantDex, err error) {
	if err = app.GameBootCamp.DB.Table("plant_dex").Select("*").Where("plant_type = ? AND is_special_item = ?", fruitType, isSpecialItem).Order("hour_to_grow desc").Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) updateGrowUpInFarm(farmId, plantDexId int) (err error) {
	if err = app.GameBootCamp.DB.Table("farm").Where("id = ?", farmId).Update(map[string]interface{}{
		"plant_dex_id": plantDexId,
		"is_watered":   false,
		"update_date":  time.Now(),
		"plant_date":   time.Now().Truncate(24 * time.Hour),
	}).Error; err != nil {
		return
	}

	return
}

func (repo *farmRepo) getItemPoolListByItemId(itemId int) (result []itemPool, err error) {
	if err = app.GameBootCamp.DB.Table("item_pool").Select("*").Where("item_id = ? AND character_id = 0", itemId).Order("id").Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) updateItemPoolById(id, itemId, characterId int) (err error) {
	if err = app.GameBootCamp.DB.Table("item_pool").Where("id = ? AND item_id = ? AND character_id = 0", id, itemId).Update(map[string]interface{}{
		"character_id": characterId,
		"update_date":  time.Now(),
	}).Error; err != nil {
		return
	}

	return
}

func (repo *farmRepo) checkRemaining(characterId, x, y int) (result remaining, err error) {
	if err = app.GameBootCamp.DB.Table("farm").Select("remaining_harvest").Where("character_id = ? AND check_point_x = ? AND check_point_y = ?", characterId, x, y).
		Find(&result).Error; err != nil {
		return
	}
	return
}


func (repo *farmRepo)checkInventory(characterId , itemId int) (result quantity,err error) {
	if err = app.GameBootCamp.DB.Table("inventory").Select("quantity").Where("character_id = ? AND item_id = ?",characterId ,itemId).

		Find(&result).Error; err != nil {
		return
	}
	return
}
func (repo *farmRepo) checkSeed(itemId int) (result stateId, err error) {
	if err = app.GameBootCamp.DB.Table("plant_dex").Select("state_id").Where(" item_id = ?", itemId).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) updateFarm(characterId, x, y, idPlantDex, remaining int) (result messageResponse, err error) {
	if err = app.GameBootCamp.DB.Table("farm").Where(" character_id = ?  AND check_point_x = ? AND check_point_y = ?  ", characterId, x, y).Update(map[string]interface{}{
		"character_id":      characterId,
		"plant_dex_id":      idPlantDex,
		"plant_date":        time.Now(),
		"update_date":       time.Now(),
		"create_date":       time.Now(),
		"remaining_harvest": remaining,
	}).Error; err != nil {
		return
	}
	result.Status = http.StatusOK
	result.MessageCode = "0000"
	result.MessageDescription = "create create username and password success"

	return
}

func (repo *farmRepo) deleteInventTory(characterId, itemId int) (result messageResponse, err error) {
	if err = app.GameBootCamp.DB.Table("inventory").Where(" character_id = ? AND item_id = ?  ", characterId, itemId).Delete(map[string]interface{}{}).Error; err != nil {
		return
	}

	result.Status = http.StatusOK
	result.MessageCode = "0000"
	result.MessageDescription = "create create username and password success"

	return
}

func (repo *farmRepo)updateQuantity(characterId , itemId,deQuantity int) (result messageResponse,err error) {
	if err = app.GameBootCamp.DB.Table("inventory").Where(" character_id = ? AND item_id = ?   ",characterId, itemId).Update(map[string]interface{}{
		"quantity": deQuantity,

	}).Error; err != nil {
		return
	}
	result.Status = http.StatusOK
	result.MessageCode = "0000"
	result.MessageDescription = "create create username and password success"

	return
}

func (repo *farmRepo) getPlantDexByitemId(itemId int) (result plantDex, err error) {
	if err = app.GameBootCamp.DB.Table("plant_dex").Select("*").Where("item_id =  ?", itemId).
		Find(&result).Error; err != nil {
		return
	}
	return
}



func (repo *farmRepo)getPlantDexByItemId(itemId int) (result plantDex, err error) {
	if err = app.GameBootCamp.DB.Table("plant_dex").Select("*").Where("item_id =  ?",itemId).

		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *farmRepo) updateWatering(characterId, checkpointX, checkpointY int, plantDexId int) (msg responseWatering, err error) {
	if err = app.GameBootCamp.DB.Table("bootcamp.dbo.farm").
		Where("character_id = ? AND check_point_x = ? AND check_point_y = ? AND plant_dex_id = ? ",
			characterId, checkpointX, checkpointY, plantDexId).
		Update(map[string]interface{}{
			"is_watered":  true,
			"update_date": time.Now(),
		}).Error; err != nil {
		return
	}
	msg.Status = http.StatusOK
	msg.MessageCode = "0000"
	msg.MessageDescription = "The plant are already watered."
	return
}
func (repo *farmRepo) getFarmWatering(characterId, checkpointX, checkpointY int) (result farm, err error) {
	if err = app.GameBootCamp.DB.Table("bootcamp.dbo.farm").
		Where("character_id = ? AND check_point_x = ? AND check_point_y = ? ",
			characterId, checkpointX, checkpointY).
		Find(&result).Error; err != nil {
		return
	}
	return
}

