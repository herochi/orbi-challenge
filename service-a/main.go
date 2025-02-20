package main

import (
	"fmt"
	"log"
	"os"

	"github/herochi/orbi/service-a/adapter/container"
	"github/herochi/orbi/service-a/config"
	"github/herochi/orbi/service-a/infrastructure/http/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var ENV = "DEV"

func main() {

	fmt.Printf("DEPLOYING SERVICE A 🚀 \n")
	config.Init()
	ENV = viper.GetString("ENV")
	log.Println("ENV: ", ENV)

	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func runServer() error {
	port := viper.GetString("PORT")
	log.Println("PORT: ", port)
	if ENV == "DEV" {
		port = viper.GetString("PORT_DEV")
	}

	g := gin.Default()
	container := container.Inject()
	server := server.NewServer(g, container)
	server.MapRoutes()

	return server.Start(port)
}
