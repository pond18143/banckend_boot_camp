package farm

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"go-api-game-boot-camp/app"

	"github.com/jinzhu/gorm"
)

var srv *farmService

type farmService struct {
	conf *app.Configs
	em   *app.ErrorMessage
	repo *farmRepo
}

func Init(conf *app.Configs, em *app.ErrorMessage) {
	srv = &farmService{
		conf: conf,
		em:   em,
		repo: &farmRepo{conf: conf},
	}
}

func NewFarmService(conf *app.Configs, em *app.ErrorMessage) *farmService {
	repo := farmRepo{}

	return &farmService{
		conf: conf,
		em:   em,
		repo: repo.InitFarmRepo(conf, em),
	}
}

func (srv *farmService) checkHeartbeat() (result heartbeatModel, err error) {

	result.Message = "farm"
	result.DateTime = time.Now()

	return
}

func (srv *farmService) getFarmInfo(characterId int) (result []loadPlantStateInFarm, err error) {

	//get Farm info by character_id
	log.Infoln("get Farm info by character_id : ", characterId)
	var farmInfo []farm
	farmInfo, err = srv.repo.getFarmInfoListByCharacterId(characterId)
	if err != nil {
		log.Errorln("srv.repo.getFarmInfoListByCharacterId fail : ", err)
		err = srv.em.Farm.GetFarmNotFound
		log.Errorln("srv.repo.getFarmInfoListByCharacterId fail : ", err)
		return
	}

	//loop plant in farm
	log.Infoln("loop plant in farm : ")
	for _, value := range farmInfo {

		//if farm not has plant next loop
		log.Infoln("if farm not has plant next loop : ", value.PlantDexId)
		if value.PlantDexId == 0 {
			result = append(result, loadPlantStateInFarm{
				FarmId:           value.Id,
				CheckPointX:      value.CheckPointX,
				CheckPointY:      value.CheckPointY,
				HarvestDate:      value.HarvestDate,
				RemainingHarvest: value.RemainingHarvest,
				PlantDate:        value.PlantDate,
				IsWatered:        value.IsWatered,
			})
			continue
		}

		//  get plant dex info by plant_dex_id
		log.Infoln("get plant dex info by plant_dex_id : ", value.PlantDexId)
		var currentPlantState plantDex
		currentPlantState, err = srv.repo.getPlantDexById(value.PlantDexId)
		if err != nil {
			log.Errorln("srv.repo.getPlantDexById fail : ", err)
			err = srv.em.Farm.GetPlantDexNotFound
			log.Errorln("srv.repo.getPlantDexById fail : ", err)
			return
		}
		//  get plant_dex in group
		log.Infoln("get plant_dex in group isSpecialItem : ", currentPlantState.PlantType)
		var plantDexState []plantDex
		isSpecialItem := false
		plantDexState, err = srv.repo.getPlantDexListByPlantType(currentPlantState.PlantType, nil, nil, &isSpecialItem)
		if err != nil {
			log.Errorln("srv.repo.getPlantDexListByPlantType fail : ", err)
			err = srv.em.Farm.GetPlantDexNotFound
			log.Errorln("srv.repo.getPlantDexListByPlantType fail : ", err)
			return
		}

		//calculate grow up for plant to next state in plant type
		log.Infoln("calculate grow up for plant to next state in plant type")

		//   plant_date diff time.now -> Day
		currentTime := time.Now()
		elapsed := currentTime.Truncate(24*time.Hour).Sub(value.PlantDate.Truncate(24*time.Hour)).Hours() / 24
		log.Infoln("plant_date diff time.now -> Day")
		log.Infoln("currentTime : ", currentTime.Truncate(24*time.Hour))
		log.Infoln("plantDate : ", value.PlantDate.Truncate(24*time.Hour))
		log.Infoln("elapsed(Day) : ", elapsed)

		plantState := loadPlantStateInFarm{
			FarmId:           value.Id,
			CheckPointX:      value.CheckPointX,
			CheckPointY:      value.CheckPointY,
			HarvestDate:      value.HarvestDate,
			RemainingHarvest: value.RemainingHarvest,
			PlantDate:        value.PlantDate,
			ElapsedTime:      elapsed,
			PlantDexId:       currentPlantState.Id,
			ItemId:           currentPlantState.ItemId,
			PlantName:        currentPlantState.PlantName,
			PlantDescription: currentPlantState.PlantDescription,
			StateId:          currentPlantState.StateId,
			StateName:        currentPlantState.StateName,
			IsWatered:        value.IsWatered,
		}

		//  check watered To Grow Up if Ripened is max Grow Up
		log.Infoln("check watered To Grow Up if Ripened is max Grow Up")
		if value.IsWatered && currentPlantState.StateId != int(Ripened) || value.IsWatered && (currentPlantState.PlantType == ForestSilk && currentPlantState.StateId != 4) {
			//   compare H <-> plant_dex.timeToGrow
			for _, growUp := range plantDexState {
				//Grow Up next start
				if growUp.PlantType == Apricot && currentPlantState.StateId == int(Seed) && elapsed >= 1 {
					//check "Apricot" if grow up -> random color
					log.Infoln("check Apricot if grow up -> random color")

					activeRandom := true
					for activeRandom {
						var randomResult int64
						randomResult, err = cryptoRandom(100)
						log.Infoln("randomResult : ", randomResult)

						if err != nil {
							log.Errorln("cryptoRandom fail : ", err)
							return
						}
						if randomResult <= int64(ApricotWhite) {
							setApricotState(srv, ApricotWhite, &plantState)
						} else if randomResult <= int64(ApricotBlue) {
							setApricotState(srv, ApricotBlue, &plantState)
						} else if randomResult <= int64(ApricotGreen) {
							setApricotState(srv, ApricotGreen, &plantState)
						} else if randomResult <= int64(ApricotRed) {
							setApricotState(srv, ApricotRed, &plantState)
						} else if randomResult <= int64(ApricotOrange) {
							setApricotState(srv, ApricotOrange, &plantState)
						} else if randomResult <= int64(ApricotPink) {
							setApricotState(srv, ApricotPink, &plantState)
						} else if randomResult <= int64(ApricotBlack) {
							setApricotState(srv, ApricotBlack, &plantState)
						} else { //ApricotRainbow
							//check item_pool
							var ApricotRainbowPool []itemPool
							ApricotRainbowPool, err = srv.repo.getItemPoolListByItemId(27)
							if err != nil {
								log.Errorln("srv.repo.getItemPoolListByItemId fail : ", err)
								err = srv.em.Farm.GetItemNotFound
								log.Errorln("srv.repo.getItemPoolListByItemId fail : ", err)
								return
							}

							//if ApricotRainbowPool is over max 10 again random
							if len(ApricotRainbowPool) == 0 {
								continue
							}

							//update ApricotRainbow pool
							setApricotState(srv, ApricotRainbow, &plantState)
							err = srv.repo.updateItemPoolById(ApricotRainbowPool[0].Id, ApricotRainbowPool[0].ItemId, characterId)
							if err != nil {
								log.Errorln("srv.repo.updateItemPoolById fail : ", err)
								err = srv.em.Farm.UpdateItemPoolFail
								log.Errorln("srv.repo.updateItemPoolById fail : ", err)
								return
							}
						}
						activeRandom = false
					} // END LOOP Random
					err = srv.repo.updateGrowUpInFarm(plantState.FarmId, plantState.PlantDexId)
					if err != nil {
						log.Errorln("srv.repo.updateGrowUpInFarm fail : ", err)
						err = srv.em.Farm.UpdateGrowUpFail
						log.Errorln("srv.repo.updateGrowUpInFarm fail : ", err)
						return
					}
					plantState.IsWatered = false
					break
				} else if currentPlantState.StateId+1 == growUp.StateId && elapsed >= float64(growUp.HourToGrow) {
					plantState.ItemId = growUp.ItemId
					plantState.PlantDexId = growUp.Id
					plantState.PlantName = growUp.PlantName
					plantState.PlantDescription = growUp.PlantDescription
					plantState.StateName = growUp.StateName
					plantState.StateId = growUp.StateId
					err = srv.repo.updateGrowUpInFarm(plantState.FarmId, plantState.PlantDexId)
					if err != nil {
						log.Errorln("srv.repo.updateGrowUpInFarm fail : ", err)
						err = srv.em.Farm.UpdateGrowUpFail
						log.Errorln("srv.repo.updateGrowUpInFarm fail : ", err)
						return
					}
					plantState.IsWatered = false
					break
				}
			} //IF compare //END LOOP compare H <-> plant_dex.timeToGrow

		} //IF IsWatered

		//check "jack o'lantern RGB" in goldenTime
		log.Infoln("check jack o'lantern RGB in goldenTime")
		if currentPlantState.PlantType == JackOLanternRGB && currentPlantState.StateId == int(Ripened) {
			if isInTimeRange(time.Kitchen, "10:00PM", "11.59:00PM") || isInTimeRange(time.Kitchen, "00:00AM", "02:00AM") {
				var goldenTimeJackOLanternRGB []plantDex
				isSpecialItem := true
				JackOLanternRGBState := _mapStateStr[Ripened]
				goldenTimeJackOLanternRGB, err = srv.repo.getPlantDexListByPlantType(JackOLanternRGB, &JackOLanternRGBState, nil, &isSpecialItem)
				if err != nil {
					log.Errorln("srv.repo.getPlantDexListByPlantType fail : ", err)
					err = srv.em.Farm.GetPlantDexNotFound
					log.Errorln("srv.repo.getPlantDexListByPlantType fail : ", err)
					return
				}

				for _, value := range goldenTimeJackOLanternRGB {
					plantState.ItemId = value.ItemId
					plantState.PlantDexId = value.Id
					plantState.PlantName = value.PlantName
					plantState.StateId = value.StateId
					plantState.PlantDescription = value.PlantDescription
				}
				log.Infoln("RGB : active mode")
			} else {
				var goldenTimeJackOLanternRGB []plantDex
				isSpecialItem := false
				JackOLanternRGBState := _mapStateStr[Ripened]
				goldenTimeJackOLanternRGB, err = srv.repo.getPlantDexListByPlantType(JackOLanternRGB, &JackOLanternRGBState, nil, &isSpecialItem)
				if err != nil {
					log.Errorln("srv.repo.getPlantDexListByPlantType fail : ", err)
					err = srv.em.Farm.GetPlantDexNotFound
					log.Errorln("srv.repo.getPlantDexListByPlantType fail : ", err)
					return
				}

				for _, value := range goldenTimeJackOLanternRGB {
					plantState.ItemId = value.ItemId
					plantState.PlantDexId = value.Id
					plantState.PlantName = value.PlantName
					plantState.StateId = value.StateId
					plantState.PlantDescription = value.PlantDescription
				}
				log.Infoln("RGB : zZzZ")

			}

			err = srv.repo.updateGrowUpInFarm(plantState.FarmId, plantState.PlantDexId)
			if err != nil {
				log.Errorln("srv.repo.updateGrowUpInFarm fail : ", err)
				err = srv.em.Farm.UpdateGrowUpFail
				log.Errorln("srv.repo.updateGrowUpInFarm fail : ", err)
				return
			}
			plantState.IsWatered = false
		}
		//add array to response
		result = append(result, plantState)
	} //END LOOP plant in farm

	return
}

func (srv *farmService) harvest(checkpointX, checkpointY, characterId int) (result harvestedResponse, err error) {

	//reload grow up in farm
	_, err = srv.getFarmInfo(characterId)
	if err != nil {
		log.Errorln("srv.getFarmInfo fail : ", err)
		err = srv.em.Farm.GetFarmNotFound
		log.Errorln("srv.getFarmInfo fail : ", err)
		return
	}

	//validate characterId AND checkpointX checkpointY
	log.Infoln("validate characterId AND checkpointX checkpointY")
	var farm farm
	farm, err = srv.repo.getFarmInfoByXY(checkpointX, checkpointY, characterId)
	if err != nil {
		log.Errorln("srv.repo.getFarmInfoByXY fail : ", err)
		err = srv.em.Farm.GetFarmNotFound
		log.Errorln("srv.repo.getFarmInfoByXY fail : ", err)
		return
	}

	//validate state can harvest RemainingHarvest >= 1
	log.Infoln("validate state can harvest RemainingHarvest : ", farm.RemainingHarvest)
	if farm.RemainingHarvest <= 0 {
		err = srv.em.Farm.HarvestFailInvalidHarvestRemaining
		log.Errorln("validate remainingHarvest fail : ", err)
		return
	}

	var currentPlantState plantDex
	currentPlantState, err = srv.repo.getPlantDexById(farm.PlantDexId)
	if err != nil {
		log.Errorln("srv.repo.getPlantDexById fail : ", err)
		err = srv.em.Farm.UpdateGrowUpFail
		log.Errorln("srv.repo.getPlantDexById fail : ", err)
		return
	}

	//validate state can harvest == Ripened
	log.Infof("validate is Apricot state id = 2 : [%+v] %d", currentPlantState.PlantType, currentPlantState.StateId)
	if (currentPlantState.PlantType == Apricot && currentPlantState.PlantType == HornyFruit) && currentPlantState.StateId != int(Ripened)-1 {
		err = srv.em.Farm.HarvestFailInvalidState
		log.Errorln("validate state fail : ", err)
		return
	}

	log.Infoln("validate state can harvest == Ripened : ", currentPlantState.StateId)
	if (currentPlantState.PlantType != Apricot && currentPlantState.PlantType != HornyFruit) && currentPlantState.StateId != int(Ripened) {
		err = srv.em.Farm.HarvestFailInvalidState
		log.Errorln("validate state fail : ", err)
		return
	}

	//validate if horny limit harvest per day
	log.Infoln("validate if horny limit harvest per day : ", currentPlantState.StateId)
	log.Infoln("plant_type : ", currentPlantState.PlantType)
	log.Infoln("plant_dax : ", currentPlantState.Id)
	log.Infoln("remaining_harvest : ", farm.RemainingHarvest)

	if currentPlantState.PlantType == HornyFruit {
		//   plant_date diff time.now -> Day
		currentTime := time.Now()
		elapsed := currentTime.Truncate(24*time.Hour).Sub(farm.PlantDate.Truncate(24*time.Hour)).Hours() / 24
		log.Infoln("plant_date diff time.now -> Day")
		log.Infoln("currentTime : ", currentTime.Truncate(24*time.Hour))
		log.Infoln("plantDate : ", farm.PlantDate.Truncate(24*time.Hour))
		log.Infoln("elapsed(Day) : ", elapsed)

		//HornyFruit next dey 1 rh=3
		//HornyFruit next dey 1 rh=2
		//HornyFruit next dey 1 rh=1
		if elapsed < 1 {
			err = srv.em.Farm.HarvestHornyFailInvalidState
			log.Errorln("validate state fail : ", err)
			return
		}
	}

	//get inventoryByCharacterId+ItemId validate item in inventory is already exist
	log.Infoln("get inventoryByCharacterId+ItemId validate item in inventory is already exist")
	ItemIsAlreadyExist := true
	var selectInventory inventory
	selectInventory, err = srv.repo.getInventoryByCharacterIdAndItemId(currentPlantState.ItemId, characterId)
	if err != nil && err == gorm.ErrRecordNotFound {
		err = nil
		ItemIsAlreadyExist = false
	} else if err != nil {
		log.Errorln("srv.repo.getInventoryByCharacterIdAndItemId : ", err)
		err = srv.em.Farm.GetInvInventoryNotFound
		log.Errorln("srv.repo.getInventoryByCharacterIdAndItemId : ", err)
		return
	}

	//open tx
	tx := app.GameBootCamp.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	//update farm if count_harvest=0 reset farm
	log.Infoln("update farm if count_harvest=0 reset farm")
	err = srv.repo.updateHarvestedInFarm(tx, checkpointX, checkpointY, farm.RemainingHarvest-1, characterId)
	if err != nil {
		log.Errorln("srv.repo.updateHarvestedInFarm : ", err)
		err = srv.em.Farm.UpdateFarmFail
		log.Errorln("srv.repo.updateHarvestedInFarm : ", err)
		tx.Rollback()
		return
	}

	//update inventory
	//validate item in inventory is already exist
	log.Infoln("validate item in inventory is already exist")
	if ItemIsAlreadyExist {
		//if is already exist "T" Update qty + 1
		err = srv.repo.updateInventoryByCharacterId(tx, currentPlantState.ItemId, selectInventory.Quantity+1, characterId)
		if err != nil {
			log.Errorln("srv.repo.updateInventoryByCharacterId : ", err)
			err = srv.em.Farm.UpdateInventoryFail
			log.Errorln("srv.repo.updateInventoryByCharacterId : ", err)
			tx.Rollback()
			return
		}
	} else {
		//if is already exist "F" NewItem
		addItem := inventory{
			ItemId:      currentPlantState.ItemId,
			Quantity:    1,
			CharacterId: characterId,
			CreateDate:  time.Now(),
			UpdateDate:  time.Now(),
		}
		err = srv.repo.addNewItemToInventoryByCharacterId(tx, addItem)
		if err != nil {
			log.Errorln("srv.repo.addNewItemToInventoryByCharacterId : ", err)
			err = srv.em.Farm.AddNewItemFail
			log.Errorln("srv.repo.addNewItemToInventoryByCharacterId : ", err)
			tx.Rollback()
			return
		}
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return
	}

	result = harvestedResponse{
		Status:             http.StatusCreated,
		MessageCode:        "0000",
		MessageDescription: "harvested successfully.",
		ItemId:             currentPlantState.ItemId,
		PlantName:          currentPlantState.PlantName,
		PlantDescription:   currentPlantState.PlantDescription,
	}

	return
}

func stringToTime(str string) time.Time {
	tm, err := time.Parse(time.Kitchen, str)
	if err != nil {
		fmt.Println("Failed to decode time:", err)
	}
	fmt.Println("Time decoded:", tm)
	return tm
}

func isInTimeRange(format string, startTimeString, endTimeString string) bool {

	t := time.Now()

	zone, offset := t.Zone()

	fmt.Println(t.Format(format), "Zone:", zone, "Offset UTC:", offset)

	timeNowString := t.Format(format)

	fmt.Println("String Time Now: ", timeNowString)

	timeNow := stringToTime(timeNowString)

	start := stringToTime(startTimeString)

	end := stringToTime(endTimeString)

	fmt.Println("Local Time Now: ", timeNow)

	if timeNow.Before(start) {
		fmt.Println("false")
		return false
	}

	if timeNow.Before(end) {
		fmt.Println("true:")
		return true
	}

	fmt.Println("false")
	return false
}

func cryptoRandom(maxNumber int64) (result int64, err error) {
	//crypto random 0-N
	var nBig *big.Int
	nBig, err = rand.Int(rand.Reader, big.NewInt(maxNumber))
	if err != nil {
		return
	}
	result = nBig.Int64() + 1
	//fmt.Printf("cryptoRandom (0-%d): %d\n", maxNumber, result)
	return
}

func setApricotState(srv *farmService, apricotColor ApricotColor, plantState *loadPlantStateInFarm) (err error) {

	plantState.PlantName = _mapRateDropApricotColor[apricotColor]
	isSpecialItem := true
	var apricotWhite []plantDex
	apricotWhite, err = srv.repo.getPlantDexListByPlantType(Apricot, nil, &plantState.PlantName, &isSpecialItem)
	if err != nil {
		return
	}

	for _, value := range apricotWhite {
		plantState.PlantDexId = value.Id
		plantState.ItemId = value.ItemId
		plantState.StateName = value.StateName
		plantState.StateId = value.StateId
		plantState.PlantDescription = value.PlantDescription
	}

	log.Infoln("randomResult : ", plantState.PlantName)

	return
}

func (srv *farmService) planting(characterId int, itemId int, x int, y int) (result messageResponse, err error) {
	//get Farm info by character_id
	remainCheck, err := srv.repo.checkRemaining(characterId, x, y) //remaining_harvest check
	log.Infof("checkRemaining :%v", err)
	if err != nil {
		log.Errorln("srv.repo.checkRemaining : ", err)
		err = srv.em.Farm.CheckRemainingNotFound
		log.Errorln("srv.repo.checkRemaining : ", err)
		return
	}
	if remainCheck.RemainingHarvest != 0 {
		log.Error("Cannot Planting Already Plant in space")

		log.Errorln("srv.repo.checkRemaining : ", err)
		err = srv.em.Farm.CheckRemainingHavePlant
		log.Errorln("srv.repo.checkRemaining : ", err)
		return
	}

	quantityCheck, err := srv.repo.checkInventory(characterId, itemId) //quentity check
	if err != nil {
		log.Errorln("srv.repo.checkInventory : ", err)
		err = srv.em.Farm.CheckInventoryNotFound
		log.Errorln("srv.repo.checkInventory : ", err)
		return
	}

	log.Infof("checkInventory :")

	log.Infof("quantity:%v", quantityCheck)
	if (quantityCheck.Quantity == 0) || (err != nil) {
		log.Error("Have No Seed")

		log.Errorln("srv.repo.checkInventory : ", err)
		err = srv.em.Farm.CheckInventoryHaveNoSeed
		log.Errorln("srv.repo.checkInventory : ", err)
		return
	}

	seedCheck, err := srv.repo.checkSeed(itemId) //state_id check
	if err != nil {
		log.Errorln("srv.repo.checkSeed : ", err)
		err = srv.em.Farm.CheckSeedNotFound
		log.Errorln("srv.repo.checkSeed : ", err)
		return
	}
	log.Infof("checkSeed :")

	if (seedCheck.StateId != 1) || (err != nil && err.Error() != "record not found") {

		log.Error("only seed can planting")
		if err != nil {
			log.Errorln("srv.repo.checkSeed : ", err)
			err = srv.em.Farm.CheckSeedOnlySeedPlant
			log.Errorln("srv.repo.checkSeed : ", err)
			return
		}
		return
	}
	PlantDexInfo, err := srv.repo.getPlantDexByItemId(itemId)
	if err != nil {
		log.Errorln("srv.repo.getPlantDexByitemId : ", err)
		err = srv.em.Farm.GetPlantDexByItemIdNotFound
		log.Errorln("srv.repo.getPlantDexByitemId : ", err)
		return
	}
	log.Infof("getPlantDexId :")
	deQuantity := quantityCheck.Quantity - 1
	if quantityCheck.Quantity == 1 {
		_, err = srv.repo.updateFarm(characterId, x, y, PlantDexInfo.Id, PlantDexInfo.Harvest)
		if err != nil {
			log.Errorln("srv.repo.addNewItemToInventoryByCharacterId : ", err)
			err = srv.em.Farm.UpdateFarmFail
			log.Errorln("srv.repo.addNewItemToInventoryByCharacterId : ", err)
			return
		}

		log.Infof("updateFarm :")
		_, err = srv.repo.deleteInventTory(characterId, itemId)
		if err != nil {
			log.Errorln("srv.repo.deleteInventTory : ", err)
			err = srv.em.Farm.DeleteInventToryFail
			log.Errorln("srv.repo.deleteInventTory : ", err)
			return
		}
		log.Infof("deleteInventTory :")
	} else {
		_, err = srv.repo.updateFarm(characterId, x, y, PlantDexInfo.Id, PlantDexInfo.Harvest)
		if err != nil {
			log.Errorln("srv.repo.updateFarm : ", err)
			err = srv.em.Farm.UpdateFarmFail
			log.Errorln("srv.repo.updateFarm : ", err)
			return
		}
		log.Infof("updateFarm :")
		_, err = srv.repo.updateQuantity(characterId, itemId, deQuantity)
		if err != nil {
			log.Errorln("srv.repo.updateQuantity : ", err)
			err = srv.em.Farm.UpdateQuantityFail
			log.Errorln("srv.repo.updateQuantity : ", err)
			return
		}
		log.Infof("updateQuantity :")
	}

	result.Status = http.StatusOK
	result.MessageCode = "0000"
	result.MessageDescription = "planting success"

	return
}

func (srv *farmService) watering(characterId, checkpointX, checkpointY int) (msg responseWatering, err error) {
	//validate get farm by characterTd checkpointX checkpointY
	var checkFarm farm
	checkFarm, err = srv.repo.getFarmWatering(characterId, checkpointX, checkpointY)
	log.Infoln(" Check status farm. ")

	if (checkFarm == farm{} || err != nil) {
		log.Infoln(" Farm not found. ")
		err = srv.em.Farm.ValidateFail.CheckFarm
		log.Error("Watering Failed : " + err.Error())
		return
	}
	log.Infoln(" Farm status : found ")
	log.Infoln(" Check status Plant. ")

	//validate plant
	if checkFarm.PlantDexId <= 0 { // plant id = 0 is false return error
		log.Infoln(" Plant not found. ")
		err = srv.em.Farm.ValidateFail.CheckPlant
		log.Error("Watering Failed : " + err.Error())
		return
	}
	log.Infoln(" Plant status : Found ")

	//validate watered
	if checkFarm.IsWatered {
		log.Info(" Already watered. ")
		err = srv.em.Farm.ValidateFail.CheckWatered
		log.Error("Watering Failed : " + err.Error())
		return
	}
	log.Info(" Haven't watered. ")

	// check watered is false update time watering in farm
	msg, err = srv.repo.updateWatering(characterId, checkpointX, checkpointY, checkFarm.PlantDexId)
	log.Infoln(" Finished watering the plants. ")
	if err != nil {
		err = srv.em.Farm.UpdateFail.UpdateWatered
		log.Error("Watering Failed : " + err.Error())
		return
	}
	return
}
