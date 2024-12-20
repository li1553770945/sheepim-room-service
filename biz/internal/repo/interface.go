package repo

import (
	"github.com/li1553770945/sheepim-room-service/biz/infra/config"
	"github.com/redis/go-redis/v9"
)

type IRepository interface {
	CreateRoom(roomId string, initClientId string) error
	IsRoomExist(roomId string) (bool, error)
	JoinRoom(roomId, userId string) error
	GetRoomMembers(roomId string) ([]string, error)
	CheckIsInRoom(roomId string, clientId string) (bool, error)
}

type Repository struct {
	Cache              *redis.Client
	CacheExpireSeconds int64
}

func NewRepository(cache *redis.Client, cfg *config.Config) IRepository {
	return &Repository{
		Cache:              cache,
		CacheExpireSeconds: cfg.CacheConfig.ExpireSeconds,
	}
}
