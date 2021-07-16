package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"go-api-game-boot-camp/app"
	"go-api-game-boot-camp/docs"
	"go-api-game-boot-camp/handler"

	log "github.com/sirupsen/logrus"
)

func main() {
	docs.SwaggerInfo.Title = "Project API"
	docs.SwaggerInfo.Description = "This is project API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"

	portAuth := flag.String("portAuth", "4440", "portAuth number")
	portApp := flag.String("portApp", "8080", "portApp number")
	configPath := flag.String("config", "configure", "set configs path, default as: 'configure'")
	stage := flag.String("stage", "dev", "set working environment")

	flag.Parse()
	log.Infof("portAuth : %+v", *portAuth)
	log.Infof("portApp : %+v", *portApp)
	log.Infof("configPath directory : %+v", *configPath)

	//parse config
	cv := app.Configs{ConfigPath: *configPath}
	//cv.BuildVersion = BuildVersion
	//cv.BuildDate = BuildDate
	if err := cv.InitViperWithStage(*stage); err != nil {
		panic(err)
	}

	log.Println("config: %+v", cv)
	em := app.ErrorMessage{Configs: cv}
	if err := em.Init(); err != nil {
		panic(err)
	}

	//connect database
		app.InitDB(&cv)

	//start http APP server
	r := handler.Routes{} //new object
	handleRoute := r.InitTransactionRoute(&cv, &em)
	AppSrv := &http.Server{
		Addr:    fmt.Sprint(":", *portApp), //":8080"
		Handler: handleRoute,
	}

	//start http Auth server
	rAuth := handler.Routes{} //new object
	handleRouteAuth := rAuth.InitTransactionRouteAuth(&cv, &em)
	AuthSrv := &http.Server{
		Addr:    fmt.Sprint(":", *portAuth), //":4440"
		Handler: handleRouteAuth,
	}

	//start http APP server
	go func() {
		if err := AppSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("transaction listen: %s\n", err)
		} else if err != nil {
			log.Panicf("transaction listen error: %s\n", err)
		}
		log.Infof("transaction listen at: %s", *portApp)
	}()

	//start http Auth server
	go func() {
		if err := AuthSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("transaction listen: %s\n", err)
		} else if err != nil {
			log.Panicf("transaction listen error: %s\n", err)
		}
		log.Infof("transaction listen at: %s", *portAuth)
	}()


	//init scheduler, sendgrid
	s := handler.InitScheduler(&cv, &em)
	s.StartScheduler()

	//sendgrid.InitSendgrid(&cv, &em)


	//create channel wait signals
	//จับสัญญาณ ctr+C
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals // wait for SIGINT
}
