//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/li1553770945/sheepim-room-service/biz/infra/cache"
	"github.com/li1553770945/sheepim-room-service/biz/infra/config"
	"github.com/li1553770945/sheepim-room-service/biz/infra/log"
	"github.com/li1553770945/sheepim-room-service/biz/infra/rpc"
	"github.com/li1553770945/sheepim-room-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-room-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-room-service/biz/internal/service"
)

func GetContainer(env string) *Container {
	panic(wire.Build(

		//infra
		config.GetConfig,
		log.InitLog,
		trace.InitTrace,
		cache.NewCache,
		rpc.NewAuthClient,

		//repo
		repo.NewRepository,

		//service
		service.NewRoomService,

		NewContainer,
	))
}
