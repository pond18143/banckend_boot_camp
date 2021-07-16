package farm

import (
	"time"
)

type FruitType int32
type PlantState int32
type ApricotColor int32

const (
	//FruitType
	GamuGamuFruit   FruitType = 1
	ForestSilk      FruitType = 2
	HornyFruit      FruitType = 3
	JackOLanternRGB FruitType = 4
	Mandragora      FruitType = 5
	Apricot         FruitType = 6

	//Plant State
	Seed    PlantState = 1
	Sprout  PlantState = 2
	Ripened PlantState = 3

	//ApricotColor
	ApricotWhite   ApricotColor = 20
	ApricotBlue    ApricotColor = 38
	ApricotGreen   ApricotColor = 54
	ApricotRed     ApricotColor = 68
	ApricotOrange  ApricotColor = 80
	ApricotPink    ApricotColor = 90
	ApricotBlack   ApricotColor = 98
	ApricotRainbow ApricotColor = 100
)

var _mapStateStr = map[PlantState]string{
	Seed:    "seed",
	Sprout:  "sprout",
	Ripened: "ripened",
}

var _mapRateDropApricotColor = map[ApricotColor]string{
	ApricotWhite:   "apricot white fruit",   //20
	ApricotBlue:    "apricot blue fruit",    //18
	ApricotGreen:   "apricot green fruit",   //16
	ApricotRed:     "apricot red fruit",     //14
	ApricotOrange:  "apricot orange fruit",  //12
	ApricotPink:    "apricot pink fruit",    //10
	ApricotBlack:   "apricot black fruit",   //8
	ApricotRainbow: "apricot rainbow fruit", //2
}

type heartbeatModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputHeartbeat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type farm struct {
	Id               int       `json:"id"`
	CharacterId      int       `json:"character_id"`
	CheckPointX      int       `json:"check_point_x"`
	CheckPointY      int       `json:"check_point_y"`
	PlantDate        time.Time `json:"plant_date"`
	HarvestDate      time.Time `json:"harvest_date"`
	RemainingHarvest int       `json:"remaining_harvest"`
	PlantDexId       int       `json:"plant_dex_id"`
	IsWatered        bool      `json:"is_watered"`
}

type plantDex struct {
	Id               int       `json:"id"`
	ItemId           int       `json:"item_id"`
	StateId          int       `json:"state_id"`
	StateName        string    `json:"state_name"`
	HourToGrow       int       `json:"hour_to_grow"`
	PlantDescription string    `json:"plant_description"`
	PlantName        string    `json:"plant_name"`
	PlantType        FruitType `json:"plant_type"`
	Harvest          int       `json:"harvest"`
}

type loadPlantStateInFarm struct {
	FarmId           int       `json:"farm_id"`
	CheckPointX      int       `json:"check_point_x"`
	CheckPointY      int       `json:"check_point_y"`
	PlantDexId       int       `json:"plant_dex_id"`
	ItemId           int       `json:"item_id"`
	PlantName        string    `json:"plant_name"`
	PlantDescription string    `json:"plant_description"`
	HarvestDate      time.Time `json:"harvest_date"`
	RemainingHarvest int       `json:"remaining_harvest"`
	PlantDate        time.Time `json:"plant_date"`
	ElapsedTime      float64   `json:"elapsed_time"`
	StateId          int       `json:"state_id"`
	StateName        string    `json:"state_name"`
	IsWatered        bool      `json:"is_watered"`
}

type itemPool struct {
	Id          int       `json:"id"`
	CharacterId int       `json:"character_id"`
	ItemId      int       `json:"item_id"`
	RoundId     int       `json:"round_id"`
	UpdateDate  time.Time `json:"update_date"`
}

type inventory struct {
	ItemId      int       `json:"item_id"`
	Quantity    int       `json:"quantity"`
	CharacterId int       `json:"character_id"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_time"`
}

type harvestingRequest struct {
	CheckPointX int `json:"check_point_x"`
	CheckPointY int `json:"check_point_y"`
}

type harvestedResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
	ItemId             int    `json:"item_id"`
	PlantName          string `json:"plant_name"`
	PlantDescription   string `json:"plant_description"`
}

type remaining struct {
	RemainingHarvest int `json:"remaining_harvest"`
}
type quantity struct {
	Quantity int `json:"quentity"`
}
type stateId struct {
	StateId int `json:"state_id"`
}

type messageResponse struct {
	Status             int64  `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}
type plantingRequest struct {
	CharacterId int `json:"character_id"`
	ItemId      int `json:"item_id"`
	CheckPointX int `json:"check_point_x"`
	CheckPointY int `json:"check_point_y"`
}
type getFarmWatering struct {
	CheckPointX int `json:"check_point_x"`
	CheckPointY int `json:"check_point_y"`
}

type responseWatering struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}
