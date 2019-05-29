package redis

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
)

type RD struct {
	myDefault *redis.Client
}

var rd RD
var Eor error

func Init(options *redis.Options) {
	rd.myDefault = redis.NewClient(options)
	_, Eor = rd.myDefault.Ping().Result()
	if Eor != nil {
		panic(Eor)
	} else {
		logs.Info("redis connect success")
	}
}

func (RD) RedisNew() *redis.Client {
	return rd.myDefault
}

func ChangeClient(client *redis.Client) RD {
	rd.myDefault = client
	return rd
}
