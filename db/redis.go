package db

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type RedisClient struct {
	addr     string
	password string
	c *redis.Client
}

var redisClient *RedisClient

func init() {
	rds, err := InitRedis()
	if err != nil {
		log.Panicln("redis connect failed !!!")
	}
	err = rds.Close()
	if err != nil {
		log.Panicln("redis connect failed !!!")
	}
	log.Println("redis connect success...")
}

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

// lpush方法封装
func (this *RedisClient) LPush(key string,data interface{}) (int64, error) {
	ret,err := this.c.LPush(key,data).Result()
	return ret, err
}

// lrange
func (this *RedisClient) LRange(key string,start int64,stop int64) ([]string, error) {
	ret,err := this.c.LRange(key,start,stop).Result()
	return ret, err
}

// lrangeAll
func (this *RedisClient) LRangeAll(key string) ([]string, error) {
	ret,err := this.c.LRange(key,0,-1).Result()
	return ret, err
}

// ZRem
func (this *RedisClient) ZRem(key string,member string) (int64, error) {
	result, err := this.c.ZRem(key, member).Result()
	return result, err
}

// ZIncrBy
func (this *RedisClient) ZIncrBy(key string,increment float64,member string) (float64, error) {
	result, err := this.c.ZIncrBy(key, increment, member).Result()
	return result, err
}

// zRevRange
func (this *RedisClient) ZRevRange(key string,start int64, stop int64) ([]string, error) {
	result, err := this.c.ZRevRange(key, start, stop).Result()
	return result, err
}

func (this *RedisClient) ZRevRangeWithScores(key string,start int64, stop int64) ([]redis.Z, error) {
	zAll ,err := this.c.ZRevRangeWithScores(key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return zAll, err
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

// Del删除key方法封装
func (this *RedisClient) Del(key string)(bool, error) {
	v, err := this.c.Del(key).Result()
	if err != nil {
		return false, err
	}
	return v > 0, nil
}

// 发布订阅
// 发布
func (this *RedisClient) Publish(topic string,message string) (int64, error) {
	result, err := this.c.Publish(topic, message).Result()
	return result,err
}

// 订阅
func (this *RedisClient) Subscribe(topic string) (*redis.PubSub) {
	subscribe := this.c.Subscribe(topic)
	return subscribe
}
