package market

import (
	"go-api-game-boot-camp/app"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type marketRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *marketRepo) InitMarketRepo(conf *app.Configs, em *app.ErrorMessage) *marketRepo {
	return &marketRepo{
		conf: conf,
		em:   em,
	}
}

func (repo *marketRepo) getMarketById(id int) (market MarketInfo, err error) {
	if err = app.GameBootCamp.DB.
		Table("bootcamp.dbo.market").
		Where("id = ?", id).
		Find(&market).Error; err != nil {
		return
	}

	return market, nil
}

func (repo *marketRepo) getItemById(id int) (item ItemInfo, err error) {
	if err = app.GameBootCamp.DB.
		Table("item").
		Where("id = ?", id).
		Find(&item).Error; err != nil {
		return
	}

	return item, nil
}

func (repo *marketRepo) getBalanceByCharacterID(id int) (balance CharacterBalance, err error) {
	if err = app.GameBootCamp.DB.
		Table("character AS ch").
		Select("ch.id, ch.login_id, ch.gold, ch.coin").
		Where("id = ?", id).
		Find(&balance).Error; err != nil {
		return
	}

	return balance, nil
}

func (repo *marketRepo) updateBalanceByCharacterID(tx *gorm.DB, moneyType string, id int, update UpdateBalance) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}

	var sqlUpdate map[string]interface{}
	if moneyType == "gold" {
		sqlUpdate = map[string]interface{}{
			"gold":        update.Gold,
			"update_date": update.UpdateDate,
		}
	} else if moneyType == "coin" {
		sqlUpdate = map[string]interface{}{
			"coin":        update.Coin,
			"update_date": update.UpdateDate,
		}
	}

	if err = tx.
		Table("character").
		Where("id = ?", id).
		Update(sqlUpdate).Error; err != nil {
		return
	}

	return nil
}

func (repo *marketRepo) getInventory(characterId int, itemId int) (item InventoryInfo, err error) {
	if err = app.GameBootCamp.DB.
		Table("inventory").
		Where("character_id = ? AND item_id = ?", characterId, itemId).
		Find(&item).Error; err != nil {
		return
	}

	return item, nil
}

func (repo *marketRepo) updateInventoryByCharacterID(tx *gorm.DB, characterId int, itemId int, new InventoryInfo) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}

	if err = tx.
		Table("inventory").
		Where("character_id = ? AND item_id = ?", characterId, itemId).
		Update(map[string]interface{}{
			"quantity":    new.Quantity,
			"update_date": new.UpdateDate}).Error; err != nil {
		return
	}

	return nil
}

func (repo *marketRepo) insertInventory(tx *gorm.DB, item InventoryInfo) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}

	if err = tx.
		Table("inventory").
		Create(&item).Error; err != nil {
		return
	}

	return
}

func (repo *marketRepo) getLottery(lotteryNum string) (lotto LotteryInfo, err error) {
	if err = app.GameBootCamp.DB.
		Table("lottery").
		Where("lottery_number = ?", lotteryNum).
		Find(&lotto).Error; err != nil {
		return
	}

	return lotto, nil
}

func (repo *marketRepo) updateLotteryPool(tx *gorm.DB, lotteryNum string, update UpdateLottery) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}

	if err = tx.
		Table("lottery").
		Where("lottery_number = ?", lotteryNum).
		Update(map[string]interface{}{
			"character_id": update.CharacterID,
			"update_date":  update.UpdateDate}).Error; err != nil {
		return
	}

	return nil
}
func (repo *marketRepo) checkItemType(tx *gorm.DB, itemId int) (result item, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.item").
		Where("id = ?", itemId).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) checkItemQuantity(tx *gorm.DB, characterId int, itemId int) (result inventory, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.inventory").
		Where("character_id = ? AND item_id = ?", characterId, itemId).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) refreshBuff(tx *gorm.DB, characterId int, buffName string) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.buff").
		Where("character_id = ? AND buff_name = ?", characterId, buffName).
		Update(map[string]interface{}{
			"remaining":   15,
			"update_date": time.Now().Format(time.RFC3339),
		}).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) updateItemLeft(tx *gorm.DB, characterId int, itemId int, itemLeft int) (result inventory, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.inventory").
		Where("character_id = ? AND item_id = ?", characterId, itemId).
		Update(map[string]interface{}{
			"quantity":    itemLeft,
			"update_date": time.Now().Format(time.RFC3339),
		}).Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) checkBuff(tx *gorm.DB, characterId int, buffName string) (result buff, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.buff").
		Where("character_id = ? AND buff_name = ?", characterId, buffName).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) insertBuff(tx *gorm.DB, request buff) (result buff, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.buff").
		Create(&request).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) updateGoldBuff(tx *gorm.DB, characterId int, gold int) (msg messageResponse, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.character").
		Where("id = ?", characterId).
		Update(map[string]interface{}{
			"gold":        gold,
			"update_date": time.Now().Format(time.RFC3339),
		}).Error; err != nil {
		return
	}
	msg.Status = http.StatusCreated
	msg.MessageCode = "0000"
	msg.MessageDescription = "The product has been sold successfully.."
	return
}

func (repo *marketRepo) updateGold(tx *gorm.DB, characterId int, gold int) (msg messageResponse, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.character").
		Where("id = ?", characterId).
		Update(map[string]interface{}{
			"gold":        gold,
			"update_date": time.Now().Format(time.RFC3339),
		}).Error; err != nil {
		return
	}
	msg.Status = http.StatusCreated
	msg.MessageCode = "0000"
	msg.MessageDescription = "The product has been sold successfully.."
	return
}

func (repo *marketRepo) updateBuff(tx *gorm.DB, characterId int, buffName string, remaining int) (err error) {

	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.buff").
		Where("character_id = ? AND buff_name = ?", characterId, buffName).
		Update(map[string]interface{}{
			"remaining":   remaining,
			"start_date": time.Now().Format(time.RFC3339),
		    "end_date": time.Now().Add(time.Hour*730).Format(time.RFC3339),
			"update_date": time.Now(),
		}).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo)getGold(tx *gorm.DB, characterId int) (result characterInfo, err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("bootcamp.dbo.character").
	Where("id = ?", characterId).
	Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) deleteInventTory(tx *gorm.DB, characterId, itemId int) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("inventory").Where("character_id = ? AND item_id = ?  ", characterId, itemId).
	Delete(map[string]interface{}{}).Error; err != nil {
		return
	}
	return
}

func (repo *marketRepo) deleteBuff(tx *gorm.DB, characterId int, buffName string) (err error) {
	if tx == nil {
		tx = app.GameBootCamp.DB
	}
	if err = tx.Table("buff").Where("character_id = ? AND buff_name = ?  ", characterId, buffName).
	Delete(map[string]interface{}{}).Error; err != nil {
		return
	}
	return
}