// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/li1553770945/sheepim-room-service/biz/infra/cache"
	"github.com/li1553770945/sheepim-room-service/biz/infra/config"
	"github.com/li1553770945/sheepim-room-service/biz/infra/log"
	"github.com/li1553770945/sheepim-room-service/biz/infra/rpc"
	"github.com/li1553770945/sheepim-room-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-room-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-room-service/biz/internal/service"
)

// Injectors from wire.go:

func GetContainer(env string) *Container {
	configConfig := config.GetConfig(env)
	traceLogger := log.InitLog()
	traceStruct := trace.InitTrace(configConfig)
	client := cache.NewCache(configConfig)
	iRepository := repo.NewRepository(client, configConfig)
	authserviceClient := rpc.NewAuthClient(configConfig)
	iRoomService := service.NewRoomService(iRepository, authserviceClient)
	container := NewContainer(configConfig, traceLogger, traceStruct, iRoomService)
	return container
}
