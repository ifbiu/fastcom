package db

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	addr     string
	password string
	c *redis.Client
}

var redisClient *RedisClient

func InitRedis() (*RedisClient,error) {
	opts := redis.Options{
		Addr:        beego.AppConfig.String("redisHost"), //redis地址
		Password:    beego.AppConfig.String("redisPassword"), //密码
		DialTimeout: time.Second * 5, //超时时间
	}
	client := redis.NewClient(&opts) //创建连接
	if err := client.Ping().Err(); err != nil { //心跳测试
		return nil,err
	}

	redisClient = &RedisClient{ //将连接赋值给全局变量
		addr:        beego.AppConfig.String("redisHost"), //redis地址
		password:    beego.AppConfig.String("redisPassword"), //密码
		c:        client,
	}
	return redisClient,nil
}

// Get方法封装
func (this *RedisClient) Get(key string) (string, error) {
	ret, err := this.c.Get(key).Result()
	return ret, err
}

// Set方法封装
func (this *RedisClient) Set(key string, val interface{}, ttl time.Duration) error {
	err := this.c.Set(key, val, ttl).Err()
	return err
}

// Exists方法封装
func (this *RedisClient) Exists(key string) (bool, error) {
	v, err := this.c.Exists(key).Result()
	if err != nil {
		return false, err
	}

	return v > 0, nil
}

// Close关闭连接方法封装
func (this *RedisClient) Close() error {
	return this.c.Close()
}

