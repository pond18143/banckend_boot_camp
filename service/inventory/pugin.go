package inventory

import (
	"go-api-game-boot-camp/app"
)

type inventoryRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *inventoryRepo) InitInventoryRepo(conf *app.Configs, em *app.ErrorMessage) *inventoryRepo {
	return &inventoryRepo{
		conf: conf,
		em:   em,
	}
}

func inventoryDetail(id int) (result []DetailInventory, err error) {
	if err = app.GameBootCamp.DB.Select("dd.id, dd.item_id, dd.quantity,dh.price_per_unit,dh.plant_description, dh.item_name, dd.create_date, dd.update_date").
		Table("bootcamp.dbo.inventory AS dd").
		Joins("INNER JOIN bootcamp.dbo.item AS dh ON dh.id = dd.item_id").
		Where("character_id = ?", id).
		Order("id").
		Find(&result).Error; err != nil {
		return
	}
	return
}

func inventoryHeader(id int) (result headerInventory, err error) {
	if err = app.GameBootCamp.DB.Select("character_id").
		Table("bootcamp.dbo.inventory").
		Where("character_id = ?", id).
		Find(&result).Error; err != nil {
		return
	}
	return
}
