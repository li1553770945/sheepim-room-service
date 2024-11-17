package cache

import (
	"context"
	"github.com/li1553770945/sheepim-room-service/biz/infra/config"
	"github.com/redis/go-redis/v9"
)

func NewCache(conf *config.Config) *redis.Client {
	cacheConfig := conf.CacheConfig

	// 构造 Redis 配置
	options := &redis.Options{
		Addr:     cacheConfig.Endpoint,
		Password: cacheConfig.Password, // 密码（如果无密码则留空）
		DB:       cacheConfig.Database, // 数据库编号（默认 0）
	}

	// 创建 Redis 客户端
	client := redis.NewClient(options)

	// 测试连接
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic("Redis 连接失败: " + err.Error())
	}

	return client
}
