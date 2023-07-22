package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/vikasd145/inventory-keeper/internal/appliance"
	"github.com/vikasd145/inventory-keeper/internal/config"
	"log"

	"github.com/vikasd145/inventory-keeper/pkg/logger"
	"github.com/vikasd145/inventory-keeper/pkg/sql"
)

var configFile = flag.String("config", "internal/config/common_local.json", "config file")

func main() {
	flag.Parse()
	config.ReadConfig(*configFile)

	err := sql.DBConn(config.Config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer sql.DB.Close()

	sql.DB.SetLogger(logger.DBLogger)

	gin.DefaultWriter = logger.File

	router := gin.Default()
	router.POST("/appliance/view/:id", appliance.ViewAppliance)
	router.POST("/appliance/search", appliance.SearchAppliance)
	router.POST("/appliance/create", appliance.CreateAppliance)
	router.POST("/appliance/update", appliance.UpdateAppliance)
	router.POST("/appliance/view/all", appliance.ViewAllAppliance)

	err = router.Run(":" + config.Config.Common.ServerPort)
	if err != nil {
		log.Fatalf("error in starting server , err:%v", err)
	}
}
