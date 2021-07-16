package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	PriKeyPath = "certification/key.auth/bootcamp.key"
	PubKeyPath = "certification/key.auth/bootcamp.key.pub"
)

type roleLogin struct {
	LoginUuid    string    `json:"login_uuid"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	RegisterDate time.Time `json:"register_date"`
	LastLogin    time.Time `json:"last_login"`
	Section      string    `json:"section"`
	ExpSection   time.Time `json:"exp_section"`
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type character struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	LoginUuid string `json:"login_uuid"`
	//Id        int    `json:"id"`
	CharacterId int    `json:"character_id"`
	LoginId     string `json:"login_id"`
	Section     string `json:"section"`
}

type Claims struct {
	CharacterId int    `json:"character_id"`
	Section     string `json:"section"`
	jwt.StandardClaims
}

type DecodeClaims struct {
	CharacterId int    `json:"character_id"`
	Section     string `json:"section"`
}

type parseCode struct {
	Value string `json:"value"`
}

type MessageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}

type inputRePassword struct {
	UserName           string `json:"user_name"`
	NewPassWord        string `json:"new_pass_word"`
	ConfirmNewPassWord string `json:"confirm_new_pass_word"`
	Otp                string `json:"otp"`
}

type resetPassword struct {
	UserName           string `json:"user_name"`
	NewPassWord        string `json:"new_pass_word"`
	ConfirmNewPassWord string `json:"confirm_new_pass_word"`
}

type inputFirstLogin struct {
	NewUserName        string `json:"new_user_name"`
	NewPassWord        string `json:"new_pass_word"`
	ConfirmNewPassWord string `json:"confirm_new_pass_word"`
	CitizenId          string `json:"citizen_id"`
}

type inputReqOTP struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type User struct {
	Id           int       `json:"id"`
	LoginUuid    string    `json:"login_uuid"`
	UserName     string    `json:"user_name"`
	PassWord     string    `json:"pass_word"`
	Name         string    `json:"name"`
	CitizenId    string    `json:"citizen_id"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
	CompanyId    int       `json:"company_id"`
	Email        string    `json:"email"`
	OTP          string    `json:"otp"`
	IsFirstLogin bool      `json:"is_first_login"`
}

type inputRegister struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
	Email            string `json:"email"`
	Gender           int    `json:"gender"`
	SkinId           int    `json:"skin_id"`
	HatId            int    `json:"hat_id"`
	ShirtId          int    `json:"shirt_id"`
	ShoesId          int    `json:"shoes_id"`
}

type user struct {
	Id           int       `json:"id"`
	LoginUuid    string    `json:"login_uuid"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	RegisterDate time.Time `json:"register_date"`
	LastLogin    time.Time `json:"last_login"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
}

type inputCharacter struct {
	Id         int       `json:"id"`
	LoginId    int       `json:"login_id"`
	Gold       int       `json:"gold"`
	Coin       int       `json:"coin"`
	Gender     int       `json:"gender"`
	SkinId     int       `json:"skin_id"`
	HatId      int       `json:"hat_id"`
	ShirtId    int       `json:"shirt_id"`
	ShoesId    int       `json:"shoes_id"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}

type inventory struct {
	Id          int       `json:"id"`
	CharacterId int       `json:"character_id"`
	ItemId      int       `json:"item_id"`
	Quantity    int       `json:"quantity"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
}

type farm struct {
	CharacterId      int       `json:"character_id"`
	CheckPointX      int       `json:"check_point_x"`
	CheckPointY      int       `json:"check_point_y"`
	PlantDate        time.Time `json:"plant_date"`
	HarvestDate      time.Time `json:"harvest_date"`
	RemainingHarvest int       `json:"remaining_harvest"`
	PlantDexId       int       `json:"plant_dex_id"`
	IsWatered        bool      `json:"is_watered"`
	CreateDate       time.Time `json:"create_date"`
	UpdateDate       time.Time `json:"update_date"`
}
