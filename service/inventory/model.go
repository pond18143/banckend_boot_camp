package inventory

import (
	"time"
)

type DetailInventory struct {
	ID               int       `json:"id"`
	ItemID           int       `json:"item_id"`
	Quantity         int       `json:"quantity"`
	ItemName         string    `json:"item_name"`
	PlantDescription string    `json:"plant_description"`
	PricePerUnit     int       `json:"price_per_unit"`
	CreateDate       time.Time `json:"create_date"`
	UpdateDate       time.Time `json:"update_time"`
}

type headerInventory struct {
	CharacterID int `json:"character_id"`
}

type responseMessage struct {
	Status             int    `json:"status"`
	MessageDescription string `json:"message_description"`
}

type resInventory struct {
	Header headerInventory `json:"header"`
	Detail []DetailInventory `json:"detail"`
}
