package config

import (
	"encoding/json"
	"fmt"
	"github.com/vikasd145/inventory-keeper/pkg/sql"
	"io/ioutil"
	"log"
	"os"
)

var Config *AppConfig

type AppConfig struct {
	Common     *CommonConfig `json:"common"`
	DBConfig   *sql.DBConfig `json:"db_config"`
	GoogleCred *Credentials  `json:"google_cred"`
}

// Credentials which stores google ids.
type Credentials struct {
	Cid      string `json:"cid"`
	Csecret  string `json:"csecret"`
	Redirect string `json:"redirect"`
}

type CommonConfig struct {
	ServerPort string `json:"server_port"`
}

func ReadConfig(configFile string) {
	Config = &AppConfig{}
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &Config); err != nil {
		log.Fatalf("unable to marshal config data,err:%v", err)

		return
	}
	fmt.Println("config loaded ", Config)
}
