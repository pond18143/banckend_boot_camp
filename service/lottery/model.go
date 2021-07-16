package lottery

import (
	"time"
)

type heartbeatModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputHeartbeat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type lotteryList struct {
	LotteryNumber string `json:"lottery_number"`
	Status        int    `json:"status"`
	Username      string `json:"username"`
}

type lotteryOutputBuy struct {
	Status bool `json:"status"`
}

type lotteryInput struct {
	Status        int    `json:"status"`
	RoundId       int    `json:"round_id"`
	LotteryNumber string `json:"lottery_number"`
	CharacterId   int    `json:"character_id"`
	PagingIndex   int    `json:"paging_index"`
	PagingSize    int    `json:"paging_size"`
}


type lotteryCount struct{
	SumQuantity int `json:"sum_quantity"`
	Amount int `json:"amount"`
}

type lotteryOutput struct {
	Header lotteryCount `json:"header"`
	Detail []lotteryList `json:"detail"`
}

type chaIdHGamu struct {
	Id int `json:"id"`
	CharacterId int `json:"character_id"`
	Quantity int `json:"quantity"`
}

type lotteryNum struct {
	LotteryNumber string `json:"lottery_number"`
}

type exchangGamu struct {
	Status string `json:"status"`
	Description string `json:"description"`
}

type responseMessage struct {
	Status             int64  `json:"status"`
	MessageDescription string `json:"message_description"`
}

type checkx struct {
	Id int `json:"id"`
	Quantity int `json:"quantity"`
}

type InventoryInfo struct {
	ID          int       `json:"id"`
	CharacterID int       `json:"character_id"`
	ItemID      int       `json:"item_id"`
	Quantity    int       `json:"quantity"`
	UpdateDate  time.Time `json:"update_date"`
	CreateDate  time.Time `json:"create_date"`
}