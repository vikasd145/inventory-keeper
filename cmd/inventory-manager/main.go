package main

import (
	"flag"
	"github.com/vikasd145/inventory-keeper/internal/api"
	"github.com/vikasd145/inventory-keeper/internal/config"
)

func main() {
	configFile := flag.String("config", "internal/config/common_local.json", "Path to config file")
	flag.Parse()

	config.ReadConfig(*configFile)
	api.RunServer()
}
