package main

import (
	"log"
	"net/http"

	"story_writer/src/common/config"
	"story_writer/src/common/database"
	"story_writer/src/constant"
	"story_writer/src/manager"
	"story_writer/src/model/usecase"
	"story_writer/src/web"
)

func main() {
	config.Init()

	cfg := config.GetConfig()

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
