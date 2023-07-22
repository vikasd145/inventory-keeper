package main

import (
	"flag"
	"github.com/vikasd145/inventory-keeper/internal/appliance"
	"github.com/vikasd145/inventory-keeper/internal/config"
	"github.com/vikasd145/inventory-keeper/pkg/sql"
	"log"
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
	//sql.DB.AutoMigrate(&user.User{}, &company.Company{}, &code.Code{})
	//sql.DB.Table("users").AutoMigrate(&user.CreateUserModel{})
	sql.DB.AutoMigrate(&appliance.Appliance{})
}
