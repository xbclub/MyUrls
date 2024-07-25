package svc

import (
	"github.com/xbclub/MyUrls/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	RedisC *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisC: redis.MustNewRedis(c.Redis),
	}
}
