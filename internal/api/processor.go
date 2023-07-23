package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vikasd145/inventory-keeper/internal/appliance"
	"github.com/vikasd145/inventory-keeper/internal/config"
	"github.com/vikasd145/inventory-keeper/pkg/logger"
	"github.com/vikasd145/inventory-keeper/pkg/sql"
	"log"
)

func RunServer() {

	db, err := sql.DBConn(config.Config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetLogger(logger.DBLogger)

	gin.DefaultWriter = logger.File

	r := gin.Default()
	r.Use(CORS())
	appModel := &appliance.Appliance{}
	appRepo := appliance.NewApplianceModel(db, appModel)
	ap := appliance.NewApplianceProcessor(appRepo, db)

	r.GET("/appliance/view/:id", ap.ViewAppliance)
	r.POST("/appliance/search", ap.SearchAppliance)
	r.POST("/appliance/create", ap.CreateAppliance)
	r.POST("/appliance/update", ap.UpdateAppliance)
	r.GET("/appliance/view/all", ap.ViewAllAppliance)

	err = r.Run(":" + config.Config.Common.ServerPort)
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %v", err))
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
