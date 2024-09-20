package redis

import (
	"context"
	"fmt"
	"sync"

	goRedis "github.com/go-redis/redis/v8"
)

type cfg struct {
	Address  string `json:"address"`
	Port     uint16 `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

var once sync.Once

var singleInstance *goRedis.Client

func GetInstance(opt ...cfg) *goRedis.Client {
	if singleInstance == nil && len(opt) > 0 {
		once.Do(func() {
			singleInstance = initRedis(&opt[0])
		})
	}
	return singleInstance
}

func initRedis(conf *cfg) *goRedis.Client {
	rc := goRedis.NewClient(&goRedis.Options{
		Addr:        fmt.Sprintf("%s:%d", conf.Address, conf.Port), // redis服务ip:port
		Password:    conf.Password,                                 // redis的认证密码
		DB:          conf.Database,                                 // 连接的database库
		IdleTimeout: 150,                                           // 默认Idle超时时间
		PoolSize:    10,                                            // 连接池
	})
	if _, err := rc.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return rc
}
