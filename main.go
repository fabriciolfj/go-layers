package main

import (
	_ "github.com/lib/pq"
	"go-layers/config"
	"go-layers/server"
	"log"
)

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing Configuration")
	config := config.InitConfig("runners")

	log.Println("Initializing Database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing Http server")
	httpServer := server.IniHttpServer(config, dbHandler)
	httpServer.Start()
}
