package main

import (
	"api/api"
	"api/config"
	"api/logs"
	"api/pkg"

	"github.com/gin-gonic/gin"
)

func main(){
	cfg := config.LoadConfig()
	logger := logs.InitLogger()
	clients := pkg.NewClients(logger, cfg)
	
	controller := api.NewController(gin.Default())
	controller.SetUpRoutes(logger, clients)
	err := controller.StartRouter(cfg)
	if err != nil{
		panic(err)
	}
}