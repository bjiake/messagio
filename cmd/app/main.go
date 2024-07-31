package main

import (
	"log"
	"messagio/pkg/config/kafka"
	"messagio/pkg/config/pgsql"
	"messagio/pkg/di"
)

func main() {
	cfg, configErr := pgsql.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load pgsql config: ", configErr)
	}
	kafkaCfg, configErr := kafka.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load kafka config: ", configErr)
	}

	server, diErr := di.InitializeAPI(cfg, kafkaCfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
