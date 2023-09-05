package main

import (
	"go-learning/src/Client"
	"go-learning/src/Interfaces"
	"go-learning/src/Utils/Functions"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load(".env")
	if env != nil {
		panic("Error loading environment\nError: " + env.Error())
	}
	Client.RunServer(Interfaces.SystemInterface{
		CorsAllowOrigin: os.Getenv("CORS_ALLOW_ORIGIN"),
		CorsAllowHeader: os.Getenv("CORS_ALLOW_HEADER"),
		CorsAllowMethod: os.Getenv("CORR_ALLOW_METHOD"),
		AppName:         os.Getenv("APP_NAME"),
		AppPrefork:      Functions.StoB(os.Getenv("APP_PREFORK")),
		AppListenPort:   Functions.StoI(os.Getenv("APP_LISTEN_PORT")),
		AppListenHost:   os.Getenv("APP_LISTEN_HOST"),
	})
}
