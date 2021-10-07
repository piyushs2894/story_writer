package main

import (
	"net/http"
	"os"
	"story_writer/src/common/config"
	"story_writer/src/common/database"
	"story_writer/src/constant"
	"story_writer/src/manager"
	"story_writer/src/model/usecase"
	"story_writer/src/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	config.Init()

	cfg := config.GetConfig()

	if _, exists := (os.LookupEnv("VERLOOP_DEBUG")); exists {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	database.Init(cfg, constant.DriverMysql)

	usecase.Init(database.DBConnMap, cfg) // Init mysql inside too

	managerMod := manager.New()

	// Initialise Web Service and HTTP Handler
	var web = web.New(managerMod)

	web.InitHandlers()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error in starting server %+v \n", err)
	}
}
