package api

import (
	"api/api/handler"
	"api/api/middleware"
	"api/config"
	"api/pkg"
	"log/slog"
	_ "api/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	SetUpRoutes(logger *slog.Logger, client pkg.Clients)
	StartRouter(cfg config.Config) error
}

type controllerImpl struct {
	Router *gin.Engine
}

func NewController(router *gin.Engine) Controller {
	return &controllerImpl{
		Router: router,
	}
}

// @title Auth Service API3
// @version 1.0
// @description This is a sample server for Auth Service.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func (R *controllerImpl) SetUpRoutes(logger *slog.Logger, client pkg.Clients) {
	h := handler.NewHandler(logger, client)
	R.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	R.Router.Use(middleware.JWTMiddleware())

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

	twit := R.Router.Group("/twits")
	twit.POST("/createTwit", h.CreateTwit)
	twit.PUT("/updateTwit", h.UpdateTwit)
	twit.DELETE("/deleteTwit/:id", h.DeleteTwit)
	twit.GET("/getUserTwit", h.GetTwits)
	twit.GET("/getFollowerTwit", h.GetFollowerTwit)
	twit.POST("/addLike", h.AddLike)
	twit.PUT("/unlike", h.RemoveLike)
}

func (R *controllerImpl) StartRouter(cfg config.Config) error {
	return R.Router.Run(cfg.API_GATEWAY)
}
