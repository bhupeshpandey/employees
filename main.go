package main

import (
	"github.com/bhupeshpandey/employees/cache"
	"github.com/bhupeshpandey/employees/config"
	"github.com/bhupeshpandey/employees/db"
	"github.com/bhupeshpandey/employees/server"
	"log"
)

func main() {
	// TODO call your code here.
	appConfig, err := config.ReadConfig()

	if err != nil {
		log.Fatal(err)
	}

	dbInst, err := db.New(appConfig.DBConfig)

	if err != nil {
		log.Fatal(err)
	}
	s := server.New(appConfig.ServerConfig, cache.New(appConfig.CacheConfig), dbInst)

	if s == nil {
		log.Fatal("Unable to create server")
	}

	s.Start()
}
