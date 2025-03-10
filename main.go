package main

import (
	"log"
	"github.com/JerryJeager/Symptomify-Backend/config"
	"github.com/JerryJeager/Symptomify-Backend/cmd"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
	config.ConnectToRedis()
}

func main() {
	log.Println("Starting Symptomify-Backend Server")
	cmd.ExecuteApiRoutes()
}
