package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go_base_project/routes"
	"os"
)

func main() {
	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	println("App started")
	godotenv.Load()
	routes.Init()
}
