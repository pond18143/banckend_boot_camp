package market

import (
	"go-api-game-boot-camp/app"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	// "go-api-game-boot-camp/service/loger"
)

var srv *marketService

type marketService struct {
	conf *app.Configs
	em   *app.ErrorMessage
	repo *marketRepo
}

func Init(conf *app.Configs, em *app.ErrorMessage) {
	srv = &marketService{
		conf: conf,
		em:   em,
		repo: &marketRepo{conf: conf},
	}
}

func NewMarketService(conf *app.Configs, em *app.ErrorMessage) *marketService {
	repo := marketRepo{}

	return &marketService{
		conf: conf,
		em:   em,
		repo: repo.InitMarketRepo(conf, em),
	}
}

func (srv *marketService) checkHeartbeat() (result heartbeatModel, err error) {

	result.Message = "market"
	result.DateTime = time.Now()

	//err = logHeartbeat(result)
	//if err != nil {
	//	return
	//}

	return
}

func (srv *marketService)SellItemTransaction(characterId int, itemId int, quantity int) (msg messageResponse, err error) {
	//create transaction -> validation, update inventory, update character
	log.Info("[-----------------Transaction Started-----------------]")
	log.Info("Item Validation")
	tx := app.GameBootCamp.DB.Begin()

	var resultCharacter characterInfo
	log.Info("Get Character Info")
	resultCharacter, err = srv.repo.getGold(tx, characterId)
	if err != nil {
		err = srv.em.Market.InsertFail.InsertCharacter
		log.Error("SellItemTransaction Failed : " + err.Error())
		tx.Rollback()
		return
	}

	if quantity <= 0 {
		err = srv.em.Market.ValidateFail.ValidateQuantityFail
		log.Error("SellItemTransaction Failed : " + err.Error())
		tx.Rollback()
		return
	}

	//validation item type can sell
	log.Info("[Item Type Validation]")
	var resultItem item
	resultItem, err = srv.repo.checkItemType(tx, itemId)
	if resultItem.MarketId != 3 && resultItem.Id != 8 || err != nil {
		err = srv.em.Market.ValidateFail.ItemTypeCantSell
		log.Error("SellItemTransaction Failed : " + err.Error())
		tx.Rollback()
		return
	}

	//validation item quantity can sell
	log.Info("[Item Quantity Validation]")
	var resultInventory inventory
	resultInventory, err = srv.repo.checkItemQuantity(tx, characterId, itemId)
	if quantity > resultInventory.Quantity || err != nil {
		err = srv.em.Market.ValidateFail.ItemQuantityCantSell
		log.Error("SellItemTransaction Failed : " + err.Error())
		tx.Rollback()
		return
	}

	//update item quantity in inventory
	log.Info("[Item Quantity Update]")
	var itemLeft int = resultInventory.Quantity - quantity
	resultInventory, err = srv.repo.updateItemLeft(tx, characterId, itemId, itemLeft)
	if resultInventory.Quantity != itemLeft || err != nil {
		err = srv.em.Market.UpdateFail.ItemQuantityCantUpdate
		log.Error("SellItemTransaction Failed : " + err.Error())
		tx.Rollback()
		return
	}

	//validate empty quantity
	log.Info("[Validate empty Update]")
	if resultInventory.Quantity == 0 {
		err = srv.repo.deleteInventTory(tx, characterId, itemId)
		if err != nil {
			err = srv.em.Market.DeleteFail.DeleteInventoryFail
			log.Error("SellItemTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}
	}

	var resultBuff buff
	if itemId == 19 {
		//validate buff mandragora
		log.Info("[Buff Mandragora Validation]")
		resultBuff, err = srv.repo.checkBuff(tx, characterId, "mandragora")
		if (resultBuff == buff{}) || (err != nil && err == gorm.ErrRecordNotFound){
			err = nil
			var inputBuff = buff {
				CharacterId: characterId,
				BuffName: "mandragora",
			}
			inputBuff.UpdateDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			inputBuff.StartDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			inputBuff.EndDate = inputBuff.StartDate.Add(time.Hour*730)
			resultBuff, err = srv.repo.insertBuff(tx, inputBuff)
			if err != nil {
				err = srv.em.Market.InsertFail.InsertBuffFail
				log.Error("SellItemTransaction Failed : " + err.Error())
				tx.Rollback()
				return
			}
		}
	}

	//first buff mandragora
	if itemId == 19 && resultBuff.Remaining == 0 {
		log.Info("[Mandragora first buff]")
		//update gold in character normal case
		var increaseGold int = (resultItem.PricePerUnit * 1) + int(float64(resultItem.PricePerUnit * (quantity - 1)) * float64(1.2))
		msg, err = srv.repo.updateGold(tx, characterId, increaseGold)
		if err != nil {
			err = srv.em.Market.UpdateFail.ItemGoldCantUpdateNormal
			log.Error("SellItemTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}
		//refresh buff
		log.Info("[Refresh buff]")
		err = srv.repo.refreshBuff(tx, characterId, "mandragora")
		if err != nil {
			err = srv.em.Market.ValidateFail.RefreshMandragoraBuffFail
			log.Error("SellItemTransaction Failed : " + err.Error())
			tx.Rollback()
			return
		}
		tx.Commit()
		return
	}

	//check buff
	log.Info("[Check buff]")
	resultBuff, err = srv.repo.checkBuff(tx, characterId, "mandragora")
	if err != nil {
		err = srv.em.Market.ValidateFail.CheckBuffFail
		log.Error("SellItemTransaction Failed : " + err.Error())
		tx.Rollback()
		return
	}

	//request > remaining
	if resultBuff.Remaining > 0 {
		if quantity > resultBuff.Remaining {
			log.Info("[Request > Remaining]")
			var	normalQuantity int = quantity - resultBuff.Remaining
			var extraQuantity int = resultBuff.Remaining
			var goldNormal = resultItem.PricePerUnit * normalQuantity
			var goldBuff int = int(float64(1.2) * float64(resultItem.PricePerUnit) * float64(extraQuantity))
			var goldUpdate = goldNormal + goldBuff

			log.Info("[Gold update")
			msg, err = srv.repo.updateGoldBuff(tx, characterId, goldUpdate + resultCharacter.Gold)
			if err != nil {
				err = srv.em.Market.UpdateFail.ItemGoldCantUpdateC1
				log.Error("SellItemTransaction Failed : " + err.Error())
				tx.Rollback()
				return
			}

			log.Info("[Update Buff]")
			err = srv.repo.updateBuff(tx, characterId, "mandragora", 0)
			if err != nil {
				err = srv.em.Market.UpdateFail.BuffCantUpdateC2
				log.Error("SellItemTransaction Failed : " + err.Error())
				tx.Rollback()
				return
			}

			log.Info("[Check buff]")
			resultBuff, err = srv.repo.checkBuff(tx, characterId, "mandragora")
			if err != nil {
				err = srv.em.Market.ValidateFail.CheckBuffFail
				log.Error("SellItemTransaction Failed : " + err.Error())
				tx.Rollback()
				return
			}

			log.Info("[Delete buff]")
			if resultBuff.Remaining == 0 && itemId != 19 {
				err = srv.repo.deleteBuff(tx, characterId, "mandragora")
				if err != nil {
					err = srv.em.Market.DeleteFail.DeleteBuffFail
					log.Error("SellItemTransaction Failed : " + err.Error())
					tx.Rollback()
					return
				}
			}

			if itemId == 19 {
				//refresh buff
				log.Info("[Refresh buff]")
				err = srv.repo.refreshBuff(tx, characterId, "mandragora")
				if err != nil {
					err = srv.em.Market.ValidateFail.RefreshMandragoraBuffFail
					log.Error("SellItemTransaction Failed : " + err.Error())
					tx.Rollback()
					return
				}
			}
		} else {
			//request <= remaining
			log.Info("[Request <= Remaining]")
			var goldBuff float64 = float64(1.2) * float64(resultItem.PricePerUnit) * float64(quantity)
			var increaseGold int = int(goldBuff)

			log.Info("[Gold update")
			msg, err = srv.repo.updateGoldBuff(tx, characterId, increaseGold + resultCharacter.Gold)
			if err != nil {
				err = srv.em.Market.UpdateFail.ItemGoldCantUpdateC2
				log.Error("SellItemTransaction Failed : " + err.Error())
				tx.Rollback()
				return
			}

			log.Info("[Update Buff]")
			err = srv.repo.updateBuff(tx, characterId, "mandragora", resultBuff.Remaining - quantity)
			if err != nil {
				err = srv.em.Market.UpdateFail.BuffCantUpdateC2
				log.Error("SellItemTransaction Failed : " + err.Error())
			}

			log.Info("[Check buff]")
			resultBuff, err = srv.repo.checkBuff(tx, characterId, "mandragora")
			if err != nil {
				err = srv.em.Market.ValidateFail.CheckBuffFail
				log.Error("SellItemTransaction Failed : " + err.Error())
				tx.Rollback()
				return
			}

			log.Info("[Delete buff]")
			if resultBuff.Remaining == 0 && itemId != 19 {
				err = srv.repo.deleteBuff(tx, characterId, "mandragora")
				if err != nil {
					err = srv.em.Market.DeleteFail.DeleteBuffFail
					log.Error("SellItemTransaction Failed : " + err.Error())
					tx.Rollback()
					return
				}
			}
			
			if itemId == 19 {
				//refresh buff
				log.Info("[Refresh buff]")
				err = srv.repo.refreshBuff(tx, characterId, "mandragora")
				if err != nil {
					err = srv.em.Market.ValidateFail.RefreshMandragoraBuffFail
					log.Error("SellItemTransaction Failed : " + err.Error())
					tx.Rollback()
					return
				}
			}
		}
		tx.Commit()
		return
	}

	//update gold in character no buff
	log.Info("[Gold normal case Update]")
	var increaseGold int = resultItem.PricePerUnit * quantity
	msg, err = srv.repo.updateGold(tx, characterId, increaseGold + resultCharacter.Gold)
	if err != nil {
		err = srv.em.Market.UpdateFail.ItemGoldCantUpdateNormal
		log.Error("SellItemTransaction Failed : " + err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
//validate item (3) if 3 = can sell
//validate quantity
//update quanttity
//validate buff mandragora
//update money
//update date

func (srv *marketService) BuyItem(characterId int, request BuyItemRequest) (err error) {
	//repo := marketRepo{}

	log.Info("character_id: ", characterId)
	log.Info("item_id: ", request.ItemID)
	log.Info("lottery_number: ", request.LotteryNumber)
	log.Info("quantity: ", request.Quantity)

	// check quantity
	if request.Quantity <= 0 {
		log.Error("validate quantity: quantity must be integer greater than 0")
		err = srv.em.Market.QuantityBadRequest
		log.Error("validate quantity: ", err)
		return err
	}

	// check item exist and get price
	item, err := srv.repo.getItemById(request.ItemID)
	if err != nil && err.Error() != "record not found" { // other than "record not found"
		log.Error("Error while get item by id: ", err)
		err = srv.em.Market.GetItemError
		log.Error("Error while get item by id: ", err)
		return err
	}

	if err != nil && err.Error() == "record not found" { // "record not found"
		log.Error("srv.repo.getItemById: no item belong to this id.")
		err = srv.em.Market.GetItemNotFound
		log.Error("srv.repo.getItemById: ", err)
		return err
	}

	// fmt.Println(item)
	log.Info("price_per_unit: ", item.PricePerUnit)
	log.Info("total_price: ", item.PricePerUnit*request.Quantity)

	// check market_id
	if item.MarketID != 1 && item.MarketID != 2 {
		log.Error("validate market_id failed: cannot buy item from this market")
		err = srv.em.Market.MarketIDBadRequest
		log.Error("validate market_id failed: ", err)
		return err
	}

	// get character balance
	balance, err := srv.repo.getBalanceByCharacterID(characterId)
	if err != nil && err.Error() != "record not found" { // other than "record not found"
		log.Error("Error while get character by id: ", err)
		err = srv.em.Market.GetCharacterError
		log.Error("Error while get character by id: ", err)
		return err
	}

	if err != nil && err.Error() == "record not found" { // "record not found"
		log.Error("srv.repo.getBalanceByCharacterID: no character belong to this id.")
		err = srv.em.Market.GetCharacterNotFound
		log.Error("srv.repo.getBalanceByCharacterID: ", err)
		return err
	}

	// fmt.Println(balance)
	// fmt.Println("characterId: ", balance.ID)
	// fmt.Println("gold: ", balance.Gold)
	// fmt.Println("coin: ", balance.Coin)

	// open tx
	tx := app.GameBootCamp.DB.Begin()

	// check character balance and update if have enough money to buy
	// coin ซื้อของ currency = 1 ได้
	// gold ซื้อของ currency = 1, 2 ได้
	if item.PermitCurrency == 2 { // ใช้ gold ซื้อ lottery
		// check lottery_number
		if request.LotteryNumber == "" {
			log.Error("validate lottery_number: please enter lottery_number")
			err = srv.em.Market.LotteryNumberBadRequest
			log.Error("validate lottery_number: ", err)
			return err
		}

		// buy lottery set qty to 1
		request.Quantity = 1
		log.Info("lottery_quantity: ", request.Quantity)
		log.Info("lottery_total_price: ", item.PricePerUnit*request.Quantity)

		// check if already buy
		lotto, err := srv.repo.getLottery(request.LotteryNumber)
		if err != nil && err.Error() != "record not found" { // other than "record not found"
			log.Error("Error while get lottery by lottery number: ", err)
			err = srv.em.Market.GetLotteryError
			log.Error("Error while get lottery by lottery number: ", err)
			return err
		}

		if lotto.CharacterID != 0 {
			log.Error("validate lottery: this lottery number is bought already")
			err = srv.em.Market.LotteryIsBought
			log.Error("validate lottery: ", err)
			return err
		}

		// check balance
		if balance.Gold < item.PricePerUnit*request.Quantity { // gold ไม่พอ
			log.Error("validate balance: do not have enough balance to buy.")
			err = srv.em.Market.NotEnoughBalance
			log.Error("validate balance: ", err)
			return err
		} else {
			// ตัด gold
			update := UpdateBalance{
				Gold:       balance.Gold - (item.PricePerUnit * request.Quantity),
				UpdateDate: time.Now(),
			}
			err = srv.repo.updateBalanceByCharacterID(tx, "gold", balance.ID, update)
			if err != nil {
				tx.Rollback()
				log.Error("Error while update balance by character id: ", err)
				err = srv.em.Market.UpdateBalanceError
				log.Error("Error while update balance by character id: ", err)
				return err
			}
			log.Info("Update gold successfully")

			// update lottery pool
			updateLottery := UpdateLottery{
				CharacterID: balance.ID,
				UpdateDate:  time.Now(),
			}

			err = srv.repo.updateLotteryPool(tx, request.LotteryNumber, updateLottery)
			if err != nil {
				tx.Rollback()
				log.Error("Error while update lottery pool by lottery_number: ", err)
				err = srv.em.Market.UpdateLotteryPoolError
				log.Error("Error while update lottery pool by lottery_number: ", err)
				return err
			}
			log.Info("Update lottery pool successfully")
		}
	} else if item.PermitCurrency == 1 { // ใช้ได้ทั้ง gold, coin ซื้อ (จะตัด coin ก่อน)
		if balance.Coin+balance.Gold < item.PricePerUnit*request.Quantity { // เงินไม่พอทั้งคู่เลย
			log.Error("validate balance: do not have enough balance to buy.")
			err = srv.em.Market.NotEnoughBalance
			log.Error("validate balance: ", err)
			return err
		} else if balance.Coin >= item.PricePerUnit*request.Quantity { // coin พอ
			// ตัด coin
			update := UpdateBalance{
				Coin:       balance.Coin - (item.PricePerUnit * request.Quantity),
				UpdateDate: time.Now(),
			}
			err = srv.repo.updateBalanceByCharacterID(tx, "coin", balance.ID, update)
			if err != nil {
				tx.Rollback()
				log.Error("Error while update balance by character id: ", err)
				err = srv.em.Market.UpdateBalanceError
				log.Error("Error while update balance by character id: ", err)
				return err
			}
			log.Info("Update coin successfully")
		} else if balance.Coin < item.PricePerUnit*request.Quantity { // coin ไม่พอ
			// ตัด coin ก่อน
			remaining := (item.PricePerUnit * request.Quantity) - balance.Coin
			update := UpdateBalance{
				Coin:       balance.Coin - balance.Coin,
				UpdateDate: time.Now(),
			}
			err = srv.repo.updateBalanceByCharacterID(tx, "coin", balance.ID, update)
			if err != nil {
				tx.Rollback()
				log.Error("Error while update balance by character id: ", err)
				err = srv.em.Market.UpdateBalanceError
				log.Error("Error while update balance by character id: ", err)
				return err
			}
			log.Info("Update coin successfully")

			// ตัด gold
			update = UpdateBalance{
				Gold:       balance.Gold - remaining,
				UpdateDate: time.Now(),
			}
			err = srv.repo.updateBalanceByCharacterID(tx, "gold", balance.ID, update)
			if err != nil {
				tx.Rollback()
				log.Error("Error while update balance by character id: ", err)
				err = srv.em.Market.UpdateBalanceError
				log.Error("Error while update balance by character id: ", err)
				return err
			}
			log.Info("Update gold successfully")
		}
	}

	// update inventory ถ้ายังไม่เคยซื้อ - insert ถ้าเคยซื้อแล้ว - update
	// get item in inventory
	inventory, err := srv.repo.getInventory(balance.ID, item.ID)
	if err != nil && err.Error() != "record not found" { // other than "record not found"
		log.Error("Error while get item in inventory: ", err)
		err = srv.em.Market.GetItemError
		log.Error("Error while get item in inventory: ", err)
		return err
	}

	if err != nil && err.Error() == "record not found" { // "record not found"
		log.Info("no item in inventory belong to this id. Creating new...")

		// insert
		charItem := InventoryInfo{
			CharacterID: balance.ID,
			ItemID:      item.ID,
			Quantity:    request.Quantity,
			UpdateDate:  time.Now(),
			CreateDate:  time.Now(),
		}
		// fmt.Println(charItem)

		err = srv.repo.insertInventory(tx, charItem)
		if err != nil {
			tx.Rollback()
			log.Error("Error while insert item in inventory: ", err)
			err = srv.em.Market.InsertInventoryError
			log.Error("Error while insert item in inventory: ", err)
			return err
		}
		log.Info("Insert item in inventory successfully")
	} else {
		// update
		// fmt.Println(inventory)

		new := InventoryInfo{
			Quantity:   inventory.Quantity + request.Quantity,
			UpdateDate: time.Now(),
		}
		err = srv.repo.updateInventoryByCharacterID(tx, balance.ID, item.ID, new)
		if err != nil {
			tx.Rollback()
			log.Error("Error while update item in inventory: ", err)
			err = srv.em.Market.UpdateInventoryError
			log.Error("Error while update item in inventory: ", err)
			return err
		}
		log.Info("Update item in inventory successfully")
	}

	log.Info("Buy item successfully. Committing...")
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Error("Error while commit transaction: ", err)
		err = srv.em.Market.CommitError
		log.Error("Error while commit transaction: ", err)
		return err
	}

	return
}
