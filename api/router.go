package api

import (
	"api/config"
	"api/pkg"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Controller interface{
	SetUpRoutes(logger *slog.Logger, client pkg.Clients)
	StartRouter(cfg config.Config)error
}

type controllerImpl struct{
	Router *gin.Engine
}

func NewController(router *gin.Engine)Controller{
	return &controllerImpl{
		Router: router,
	}
}

func(R *controllerImpl) SetUpRoutes(logger *slog.Logger, client pkg.Clients){
	
}

func(R *controllerImpl) StartRouter(cfg config.Config)error{
	return R.Router.Run(cfg.API_GATEWAY)
}