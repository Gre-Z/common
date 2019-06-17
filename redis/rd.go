package redis

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
)

type RD struct {
	myDefault *redis.Client
}

var rd RD
var Eor error

type Options struct {
	//Network            string
	Addr     string
	Password string
	DB       int
	//MaxRetries         int
	//MinRetryBackoff    time.Duration
	//MaxRetryBackoff    time.Duration
	//DialTimeout        time.Duration
	//ReadTimeout        time.Duration
	//WriteTimeout       time.Duration
	//PoolSize           int
	//MinIdleConns       int
	//MaxConnAge         time.Duration
	//PoolTimeout        time.Duration
	//IdleTimeout        time.Duration
	//IdleCheckFrequency time.Duration
	//TLSConfig          *tls.Config
}

func Init(options Options) {
	x := redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	}
	rd.myDefault = redis.NewClient(&x)
	_, Eor = rd.myDefault.Ping().Result()
	if Eor != nil {
		panic(Eor)
	} else {
		logs.Info(fmt.Sprintf("redis[%d] 连接成功", x.DB))
	}
}

func (RD) RedisNew() *redis.Client {
	return rd.myDefault
}

func ChangeClient(client *redis.Client) RD {
	rd.myDefault = client
	return rd
}
