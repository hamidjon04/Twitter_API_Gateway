package api

import (
	"api/api/handler"
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
	h := handler.NewHandler(logger, client)

	user := R.Router.Group("/user")
	user.GET("/getUsers", h.GetUsers)
	user.DELETE("/deleteUser", h.DeleteUsers)
	user.GET("/getUserById", h.GetByIdUsers)
	user.GET("/getFollowers", h.GetFollowers)
	user.DELETE("/deleteFollower", h.DeleteFollower)
	user.GET("/getByIdFollower", h.GetByIdFollower)
	user.POST("/subscribe", h.Subscribe)
	user.GET("/getFollowing", h.GetFollowing)
	user.DELETE("/deleteFollowing", h.DeleteFollowing)
	user.GET("/getByIdFollowing", h.GetByIdFollowing)
}

func(R *controllerImpl) StartRouter(cfg config.Config)error{
	return R.Router.Run(cfg.API_GATEWAY)
}