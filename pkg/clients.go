package pkg

import (
	"api/config"
	"api/generated/twit"
	"api/generated/users"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients interface {
	UserClient() users.UserServiceClient
	TwitClient() twit.TwitServiceClientClient
}

type clientImpl struct {
	Logger *slog.Logger
	Cfg    config.Config
}

func NewClients(logger *slog.Logger, cfg config.Config) Clients {
	return &clientImpl{
		Logger: logger,
		Cfg:    cfg,
	}
}

func (C *clientImpl) UserClient() users.UserServiceClient {
	grpc, err := grpc.NewClient(C.Cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		C.Logger.Error(fmt.Sprintf("User service bilan bog'lanib bo'lmadi: %v", err))
		return nil
	}
	return users.NewUserServiceClient(grpc)
}

func (C *clientImpl) TwitClient() twit.TwitServiceClientClient{
	grpc, err := grpc.NewClient(C.Cfg.TWIT_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		C.Logger.Error(fmt.Sprintf("Twit service bilan bog'lanib bo'lmadi: %v", err))
		return nil
	}
	return twit.NewTwitServiceClientClient(grpc)
}


