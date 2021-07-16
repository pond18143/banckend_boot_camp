package lottery

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"go-api-game-boot-camp/app"
	"time"
)

type lotteryRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *lotteryRepo) InitLotteryRepo(conf *app.Configs, em *app.ErrorMessage) *lotteryRepo {
	return &lotteryRepo{
		conf: conf,
		em:   em,
	}
}

func (repo *lotteryRepo) logHeartbeat(hb heartbeatModel) (err error) { //sql
	if err = app.GameBootCamp.DB.Table("document_header").Save(hb).Error; err != nil {
		return
	}
	return
}

func LotteryAll(request lotteryInput) (result lotteryOutput, err error) {

	amount := app.GameBootCamp.DB.Select("COUNT(lottery_number) AS amount").
		Table("lottery AS lt").
		Where("round_id =?", request.RoundId)

	data := app.GameBootCamp.DB.Select("lt.lottery_number,li.username,CASE WHEN character_id = 0 THEN 1 ELSE 2 END AS status").
		Table("lottery AS lt").
		Joins("LEFT JOIN character AS ch ON lt.character_id = ch.id").
		Joins("LEFT JOIN login AS li ON ch.login_id = li.id").
		Where("lt.round_id =?", request.RoundId).
		Order("lt.lottery_number")

	//ยังไม่ขายทั้งหมด "filter":1
	if request.Status == 1 {
		data = data.Where("lt.character_id=?", 0)
		amount = amount.Where("character_id=?", 0)
	}
	//ขายแล้วทั้งหมด "filter":2
	if request.Status == 2 {
		data = data.Where("NOT lt.character_id=?", 0)
		amount = amount.Where("NOT character_id=?", 0)
	}
	//ค้นหาจาก CharacterId
	if request.CharacterId != 0 {
		data = data.Where("lt.character_id=?", request.CharacterId)
		amount = amount.Where("character_id=?", request.CharacterId)
	}
	//ค้นหาจาก LotteryNumber
	if request.LotteryNumber != "" {
		data = data.Where("lt.lottery_number=?", request.LotteryNumber)
		amount = amount.Where("lottery_number=?", request.LotteryNumber)
	}
	//เริ่มจาก row ที่
	if request.PagingIndex != 0 {
		data = data.Offset((request.PagingIndex - 1) * request.PagingSize)
	}
	//แสดงกี่ row
	if request.PagingSize != 0 {
		data = data.Limit(request.PagingSize)
	}

	if err = amount.Find(&result.Header).Error; err != nil {
		return
	}
	if err = data.Find(&result.Detail).Error; err != nil {
		return
	}

	return
}

func CheckGamu() (result []chaIdHGamu,err error){
	if err = app.GameBootCamp.DB.Select("id,character_id,quantity").
		Table("inventory").
		Where("item_id=8").
		Find(&result).Error;err != nil {
		return
	}
	return
}

func LotteryCount() (result lotteryCount,err error) {
	if err = app.GameBootCamp.DB.Select("SUM(quantity) AS sum_quantity,COUNT(id) AS amount").
		Table("inventory").
		Where("item_id=8").
		Find(&result).Error;err != nil {
		return
	}
	return
}

func RandomLottery(request int) (result []lotteryNum ,err error){
	if err = app.GameBootCamp.DB.Select("lottery_number").
		Table("lottery").
		Where("round_id =? AND character_id =?",1,0).
		Order("NEWID()").
		Limit(request).
		Find(&result).Error;err != nil {
		return
	}
	return
}

func AddCharId(CharId []chaIdHGamu,LotNum []lotteryNum,Count lotteryCount) (err error) {
	log.Info("[-----------------Transaction Started-----------------]")
	app.GameBootCamp.DB.Transaction(func(tx *gorm.DB) error {
		k:=0
		for i:=0;i<Count.Amount;i++ {
			lotteryquantity:=CharId[i].Quantity
			for ;CharId[i].Quantity>0;{
				log.Infof("[-----------------[ %d ]-----------------]",k)
				log.Infof("result: %+v", CharId[i])
				log.Infof("result: %+v", LotNum[k])
				//update lottery owner
				if err = tx.Table("bootcamp.dbo.lottery").
					Where("round_id =? AND character_id =? AND lottery_number =?", 1, 0, LotNum[k].LotteryNumber).
					Update(map[string]interface{}{
						"character_id": CharId[i].CharacterId,
						"update_date":  time.Now(),
						"buy_by":1,
					}).
					Error; err != nil {
					return err
				}
				CharId[i].Quantity--
				//Delete Gamu
				if CharId[i].Quantity==0 {
					log.Info("[-----------------[ delete and update ]-----------------]")
					if err = tx.Table("bootcamp.dbo.inventory").
						Where("id =? AND character_id = ? ", CharId[i].Id, CharId[i].CharacterId).
						Delete(map[string]interface{}{
						}).Error; err != nil {
						return err
					}
					var check checkx
					if err = tx.Select("TOP 1 id,quantity").
						Table("bootcamp.dbo.inventory").
						Where("character_id = ? AND item_id = ?", CharId[i].CharacterId,7).
						Find(&check).Error;err != nil {
						log.Info("check error")
						log.Infof("result: %+v", check)
						check.Id=0
						check.Quantity=0
						err =nil
					}
					lotteryquantity+=check.Quantity
					if check.Quantity>0 {
						log.Info("[-----------------[ update lottery in inventory ]-----------------]")
						if err = tx.Table("bootcamp.dbo.inventory").
							Where("id=? AND item_id =? AND character_id = ? ",check.Id, 7, CharId[i].CharacterId).
							Update(map[string]interface{}{
								"quantity": lotteryquantity,
								"update_date": time.Now(),
							}).
							Error; err != nil {
							return err
						}
						log.Infof("result: {Id:%d CharacterId:%d ItemId:%d Quantity:%d", check.Id,lotteryquantity,7,lotteryquantity)
					}else {
						log.Info("[-----------------[ create item lottery in inventory ]-----------------]")
						item := InventoryInfo{
							CharacterID: CharId[i].CharacterId,
							ItemID:      7,
							Quantity:    lotteryquantity,
							UpdateDate:  time.Now(),
							CreateDate:  time.Now(),
						}
						if err = tx.Table("inventory").
							Create(&item).Error; err != nil {
							return err
						}
						log.Infof("result: %+v", item)
					}
				}
				k++
			}
		}
		return nil
	})
	log.Info("[-----------------Transaction Stop-----------------]")
	return
}
