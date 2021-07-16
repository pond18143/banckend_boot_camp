package handler

import (
	"go-api-game-boot-camp/service/lottery"
	"net/http"

	"go-api-game-boot-camp/app"
	"go-api-game-boot-camp/service/auth"
	"go-api-game-boot-camp/service/character"
	"go-api-game-boot-camp/service/farm"
	"go-api-game-boot-camp/service/inventory"
	"go-api-game-boot-camp/service/market"
	"go-api-game-boot-camp/service/ping"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	Validation  gin.HandlerFunc
}

type Routes struct {
	pingService      []route
	characterService []route
	farmService      []route
	marketService    []route
	lotteryService   []route
}

func (r Routes) InitTransactionRoute(cv *app.Configs, em *app.ErrorMessage) http.Handler {
	middleware := NewMidleware(cv, em)

	ping := ping.NewEndpoint(cv, em)
	character := character.NewEndpoint(cv, em)
	farm := farm.NewEndpoint(cv, em)
	market := market.NewEndpoint(cv, em)
	lottery := lottery.NewEndpoint(cv, em)
	inventory := inventory.NewEndpoint(cv, em)

	r.pingService = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/",
			Endpoint:    ping.PingGetEndpoint,
		},
		{
			Name:        "Ping Pong : GET Prams",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/:name",
			Endpoint:    ping.PingGetParamsEndpoint,
		},
		{
			Name:        "Ping Pong : POST Prams+Body",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodPost,
			Pattern:     "/:name",
			Endpoint:    ping.PingPostParamsAndBodyEndpoint,
		},
	}

	r.characterService = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/",
			Endpoint:    character.PingGetEndpoint,
		},
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Get CharacterInfo",
			Method:      http.MethodGet,
			Pattern:     "/getCharacterInfo",
			Endpoint:    character.GetCharacterInfo,
			Validation:  middleware.ValidateRequestHeader,
		},
		{
			Name:        "Inventory : POST",
			Description: "Inventory : item list",
			Method:      http.MethodGet,
			Pattern:     "/getInventory",
			Endpoint:    inventory.GetInventory,
			Validation:  middleware.ValidateRequestHeader,
		},
	}
	r.farmService = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/",
			Endpoint:    farm.PingGetEndpoint,
			Validation:  middleware.ValidateRequestHeader,
		},
		{
			Name:        "getFarmInfo",
			Description: "get data for render farm",
			Method:      http.MethodGet,
			Pattern:     "/getFarmInfo",
			Endpoint:    farm.GetFarmInfo,
			Validation:  middleware.ValidateRequestHeader,
		}, {

			Name:        "harvest",
			Description: "get data for render farm",
			Method:      http.MethodPost,
			Pattern:     "/harvest",
			Endpoint:    farm.Harvest,
			Validation:  middleware.ValidateRequestHeader,
		}, {
			Name:        "watering",
			Description: "watered for growth of plant",
			Method:      http.MethodPost,
			Pattern:     "/watering",
			Endpoint:    farm.Watering,
			Validation:  middleware.ValidateRequestHeader,
		},
		{
			Name:        "harvest",
			Description: "get data for render farm",
			Method:      http.MethodPost,
			Pattern:     "/planting",
			Endpoint:    farm.Planting,
			Validation:  middleware.ValidateRequestHeader,
		},
	}

	r.marketService = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/",
			Endpoint:    market.PingGetEndpoint,
		},
		{
			Name:        "sellItem",
			Description: "sell item",
			Method:      http.MethodPost,
			Pattern:     "/sell",
			Endpoint:    market.SellItem,
			Validation:  middleware.ValidateRequestHeader,
		},
		{
			Name:        "Buy Item : POST",
			Description: "Buy Item : Market",
			Method:      http.MethodPost,
			Pattern:     "/buy",
			Endpoint:    market.BuyItem,
			Validation:  middleware.ValidateRequestHeader,
		},
	}

	r.lotteryService = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/",
			Endpoint:    lottery.PingGetEndpoint,
			Validation:  middleware.ValidateRequestHeader,
		}, {
			Name:        "LotteryStock",
			Description: "LotteryStock",
			Method:      http.MethodPost,
			Pattern:     "/stock",
			Endpoint:    lottery.LotteryStock,
			Validation:  middleware.ValidateRequestHeader,
		}, {
			Name:        "RandomLottery",
			Description: "RandomLottery",
			Method:      http.MethodPost,
			Pattern:     "/exchange",
			Endpoint:    lottery.RandomLottery,
			Validation:  middleware.ValidateRequestHeader,
		}, /*{
			Name:        "BuyLottery",
			Description: "BuyLottery",
			Method:      http.MethodPost,
			Pattern:     "/BuyLottery",
			Endpoint:    lottery.BuyLottery,
			Validation:  middleware.ValidateRequestHeader,
		},*/
	}

	ro := gin.New()

	ro.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "x-api-key"},
		AllowCredentials: true,
	}))

	store := ro.Group("/app/ping")
	for _, e := range r.pingService {
		store.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}

	store = ro.Group("/app/character")
	for _, e := range r.characterService {
		store.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}

	store = ro.Group("/app/farm")
	for _, e := range r.farmService {
		store.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}

	store = ro.Group("/app/market")
	for _, e := range r.marketService {
		store.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}

	store = ro.Group("/app/lottery")
	for _, e := range r.lotteryService {
		store.Handle(e.Method, e.Pattern, e.Validation, e.Endpoint)
	}
	return ro
}

func (rAuth Routes) InitTransactionRouteAuth(cv *app.Configs, em *app.ErrorMessage) http.Handler {

	iojwt := auth.NewEndpoint(cv, em)

	txiojwt := []route{
		{
			Name:        "register",
			Description: "register",
			Method:      http.MethodPost,
			Pattern:     "/register",
			Endpoint:    iojwt.Register,
		},
		{
			Name:        "signin",
			Description: "signin",
			Method:      http.MethodPost,
			Pattern:     "/signin",
			Endpoint:    iojwt.SignIn,
		},
	}

	roAuth := gin.New()
	storeAuth := roAuth.Group("/auth")
	for _, e := range txiojwt {
		storeAuth.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return roAuth
}
