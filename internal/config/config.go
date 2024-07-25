package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis          redis.RedisConf
	ShortKeyLength int `json:",default=7"`
	ShortKeyTTL    int `json:",default=604800"`
	WebSiteURL     string
}
