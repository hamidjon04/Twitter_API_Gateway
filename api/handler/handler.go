package handler

import (
	"api/generated/twit"
	"api/generated/users"
	"api/pkg"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetUsers(c *gin.Context)
	DeleteUsers(c *gin.Context)
	GetByIdUsers(c *gin.Context)
	GetFollowers(c *gin.Context)
	DeleteFollower(c *gin.Context)
	GetByIdFollower(c *gin.Context)
	Subscribe(c *gin.Context)
	GetFollowing(c *gin.Context)
	DeleteFollowing(c *gin.Context)
	GetByIdFollowing(c *gin.Context)
}

type handlerImpl struct {
	Logger      *slog.Logger
	UserService users.UserServiceClient
	TwitService twit.TwitServiceClientClient
}

func NewHandler(logger *slog.Logger, client pkg.Clients) Handler {
	return &handlerImpl{
		Logger:      logger,
		UserService: client.UserClient(),
		TwitService: client.TwitClient(),
	}
}
