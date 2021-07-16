package character

import "time"

type heartbeatModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputHeartbeat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type characterDetail struct{
    Id int `json:"id"`
    Username string `json:"username"`
	LoginId int `json:"login_id"`
    Gold int `json:"gold"`
	Coin int `json:"coin"`
	Gender int `json:"gender_id"`
	SkinId int `json:"skin_id"`
	HatId int `json:"hat_id"`
	ShirtId int `json:"shirt_id"`
	ShoesId int `json:"shoes_id"`
	UpdateDate  time.Time `json:"update_date"`
}
type buffDetail struct{
    
	BuffName string `json:"buff_name"`
	Remaining int `json:"remaining"`
	UpdateDate  time.Time `json:"update_date"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description  string   `json:"description"`
	Value     int `json:"value"`
}
//type buffDetail struct{
//	CharacterId int `json:"character"`
//	BuffName string `json:"BuffName"`
//	Remaining int `json:"remaining"`
//	UpdateDate  time.Time `json:"update_date"`
//	StartDate   time.Time `json:"start_date"`
//	EndDate     time.Time `json:"end_date"`
//}





type characterDetailRes struct{
	Character characterDetail `json:"character"`
	Buff []buffDetail `json:"buff"`
}

type messageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}
