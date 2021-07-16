package market

import "time"

type heartbeatModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputHeartbeat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type messageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}

type inputSellItem struct {
	ItemId				int			`json:"item_id"`
	Quantity			int			`json:"quantity"`
}

type item struct {
	Id               int       `json:"id"`
	MarketId         int       `json:"market_id"`
	ItemName         string    `json:"item_name"`
	PricePerUnit     int       `json:"price_per_unit"`
	ItemType         string    `json:"item_type"`
	PermitCurrency   int       `json:"permit_currency"`
	UpdateDate       time.Time `json:"update_date"`
	PlantDescription string    `json:"plant_description"`
}

type inventory struct {
	Id          int       `json:"id"`
	CharacterId int       `json:"character_id"`
	ItemId      int       `json:"item_id"`
	Quantity    int       `json:"quantity"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
}

type buff struct {
	Id          int       `json:"id"`
	CharacterId int       `json:"character_id"`
	BuffName    string    `json:"buff_name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Remaining   int       `json:"remaining"`
	Description string    `json:"description"`
	Value       int       `json:"value"`
	UpdateDate  time.Time `json:"update_date"`
}

type BuyItemRequest struct {
	ItemID        int    `json:"item_id"`
	LotteryNumber string `json:"lottery_number"`
	Quantity      int    `json:"quantity"`
}

type MarketInfo struct {
	ID         int       `json:"id"`
	MarketName string    `json:"market_name"`
	MarketDesc string    `json:"market_desc"`
	UpdateDate time.Time `json:"update_date"`
}

type ItemInfo struct {
	ID               int       `json:"id"`
	MarketID         int       `json:"market_id"`
	ItemName         string    `json:"item_name"`
	PricePerUnit     int       `json:"price_per_unit"`
	ItemType         string    `json:"item_type"`
	PermitCurrency   int       `json:"permit_currency"`
	Update_Date      time.Time `json:"update_date"`
	PlantDescription string    `json:"plant_description"`
}

type CharacterBalance struct {
	ID      int `json:"id"`
	LoginID int `json:"login_id"`
	Gold    int `json:"gold"`
	Coin    int `json:"coin"`
}

type UpdateBalance struct {
	Gold       int       `json:"gold"`
	Coin       int       `json:"coin"`
	UpdateDate time.Time `json:"update_date"`
}

type InventoryInfo struct {
	ID          int       `json:"id"`
	CharacterID int       `json:"character_id"`
	ItemID      int       `json:"item_id"`
	Quantity    int       `json:"quantity"`
	UpdateDate  time.Time `json:"update_date"`
	CreateDate  time.Time `json:"create_date"`
}

type characterInfo struct {
	Id				int			`json:"id"`
	LoginId			int			`json:"login_id"`
	Gold			int			`json:"gold"`
	Coin			int			`json:"coin"`
	Gender			int			`json:"gender"`
	SkinId			int			`json:"skin_id"`
	HatId			int			`json:"hat_id"`
	ShirtId			int			`json:"shirt_id"`
	ShoesId			int			`json:"shoes_id"`
	CreateDate		time.Time	`json:"create_date"`
	UpdateDate		time.Time	`json:"update_date"`
}
type UpdateLottery struct {
	CharacterID int       `json:"character_id"`
	UpdateDate  time.Time `json:"update_date"`
}

type LotteryInfo struct {
	ID            int       `json:"id"`
	CharacterID   int       `json:"character_id"`
	LotteryNumber string    `json:"lottery_number"`
	RoundID       int       `json:"round_id"`
	UpdateDate    time.Time `json:"update_date"`
	CreateDate    time.Time `json:"create_date"`
	BuyBy         int       `json:"buy_by"`
}
