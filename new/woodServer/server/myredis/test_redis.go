package myredis

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RRedis struct {
	redisCli       *redis.Pool
	maxIdle        int
	maxActive      int
	maxIdleTimeout time.Duration
	maxTimeout     time.Duration
	lazyLimit      bool
	maxSize        int
}

// 从池里获取连接 ———— 单独的方法
func (r *RRedis) getRedisConn() redis.Conn {
	rc := r.redisCli.Get()
	// // 用完后将连接放回连接池
	//defer rc.Close()
	return rc
}

// GetAllKeys 获取所有keys
func (r *RRedis) GetAllKeys() []string {

	rc := r.getRedisConn()
	defer rc.Close()

	keys, err := redis.Strings(rc.Do("KEYS", "*"))
	if err != nil {
		return make([]string, 0)
	}
	return keys
}

func (r *RRedis) Exists(key string) bool {
	rc := r.getRedisConn()
	flag, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return flag
}

func (r *RRedis) Get(key string, timeout int) (string, error) {

	start := time.Now()

	for {
		res, err := r.GetNoWait(key)
		if err != nil {
			return "", err
		} else if res == "" {
			if timeout != -1 {
				lasted := time.Now().Sub(start)
				if r.maxTimeout > lasted {
					t1 := r.maxTimeout
					t2 := time.Duration(timeout)*time.Second - lasted
					time.Sleep(time.Duration(math.Min(float64(t1), float64(t2))))
				} else {
					return "", errors.New("GET timeout")
				}
			} else {
				time.Sleep(r.maxTimeout)
			}
		} else {
			return res, nil
		}
	}
}

// func (r *RRedis) GetNoWait(key string) (string, error) {

// 	rc := r.getRedisConn()
// 	defer rc.Close()

// 	res, err := redis.String(rc.Do("LPOP", key))

// 	if err != nil {
// 		return "", err
// 	}
// 	return res, nil
// }

func (r *RRedis) Put(key string, value string, timeout int) (int, error) {

	start := time.Now()

	for {
		res, err := r.PutNoWait(key, value)

		if err != nil {
			return 0, err
		} else if res == -1 {
			if timeout != -1 {
				lasted := time.Now().Sub(start)
				if r.maxTimeout > lasted {
					t1 := r.maxTimeout
					t2 := time.Duration(timeout)*time.Second - lasted
					time.Sleep(time.Duration(math.Min(float64(t1), float64(t2))))
				} else {
					return 0, errors.New("PUT timeout")
				}
			} else {
				time.Sleep(r.maxTimeout)
			}

		} else {
			return res, nil
		}

	}
}

func (r *RRedis) PutNoWait(key string, value string) (int, error) {

	rc := r.getRedisConn()
	defer rc.Close()

	if r.Full(key) {
		return -1, nil
	}

	res, err := redis.Int(rc.Do("RPUSH", key, value))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) QSize(key string) int {

	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("LLEN", key))
	if err != nil {
		return -1
	}

	return res
}

func (r *RRedis) Empty(key string) bool {

	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("LLEN", key))
	if err != nil {
		return false
	}
	if res == 0 {
		return true
	}
	return false
}

func (r *RRedis) Full(key string) bool {

	if r.maxSize != 0 && r.QSize(key) >= r.maxSize {
		return true
	}

	return false
}

func (r *RRedis) GetNoWait(key string) (string, error) {

	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.String(rc.Do("LPush", key))

	if err != nil {
		return "", err
	}
	return res, nil
}

// 父级interface
type RedisFatherInterface interface {
	// 获取全部keys
	GetAllKeys() []string
}

type StringInterface interface {
	RedisFatherInterface
	StringSet(key string, value string)
	StringGet(key string) (string, error)
	StringExists(key string) bool
}

// ListInterface 操做list接口
type ListInterface interface {
	// "继承"父类的全部方法
	RedisFatherInterface
	GetNoWait(key string) (string, error) // ~di: interface{}类型、数据压缩
	Get(key string, timeout int) (string, error)
	Put(key string, value string, timeout int) (int, error)
	PutNoWait(key string, value string) (int, error)
	QSize(key string) int
	Empty(key string) bool
	Full(key string) bool
	Exists(key string) bool

	ListAdd(list string, key string)
}

type HashInterface interface {
	RedisFatherInterface
	HashAdd(hash string, key string, value string)
	HashExist(hash string, key string) bool
	HashGet(hash string, key string) (string, error)
}

type ZsetInterface interface {
	ZSetAdd(set string, key string, value float64)
}

// redis全部操做的接口
type RedisInterface interface {
	ListInterface
	HashInterface
	ZsetInterface
	StringInterface

	IsNill(error interface{}) bool
}

// 工厂函数，要求对应的结构体必须实现 RedisInterface 中的全部方法
// 若是只想实现某一些方法，就返回"有这些方法的结构体"就行了
func ProduceRedis(host, port, password string, db, maxSize int, lazyLimit bool) (RedisInterface, error) {

	// 要求RRedis结构体实现返回的接口中全部的方法！
	redisObj := &RRedis{
		maxIdle:        100,
		maxActive:      130,
		maxIdleTimeout: time.Duration(60) * time.Second,
		maxTimeout:     time.Duration(30) * time.Second,
		lazyLimit:      lazyLimit,
		maxSize:        maxSize,
	}
	// 创建链接池
	redisObj.redisCli = &redis.Pool{
		MaxIdle:     redisObj.maxIdle,
		MaxActive:   redisObj.maxActive,
		IdleTimeout: redisObj.maxIdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(
				"tcp",
				host+":"+port, // address
				redis.DialPassword(password),
				redis.DialDatabase(int(db)),
				redis.DialConnectTimeout(redisObj.maxTimeout),
				redis.DialReadTimeout(redisObj.maxTimeout),
				redis.DialWriteTimeout(redisObj.maxTimeout),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

	return redisObj, nil
}

func (r *RRedis) ZSetAdd(set string, key string, value float64) {

	rc := r.getRedisConn()
	defer rc.Close()

	_, err := rc.Do("ZADD", set, value, key)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}


func (r *RRedis) ListAdd(list string, key string) {
	rc := r.getRedisConn()
	defer rc.Close()

	_, err := rc.Do("LPUSH", list, key)
	if err != nil {
		fmt.Println("redis list failed:", err)
	}
}

func (r* RRedis) HashAdd(hash string, key string, value string) {
	rc := r.getRedisConn()
	defer rc.Close()

	_, err := rc.Do("HSET", hash, key, value)
	if err != nil {
		fmt.Println("redis list failed:", err)
	}
}

func (r* RRedis) HashGet(hash string, key string) (string, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	ret, err := redis.String(rc.Do("HGET", hash, key))
	if err != nil {
		fmt.Println("redis list failed:", err)
		return "", err
	}
	return ret, nil
}

func (r* RRedis) HashExist(hash string, key string) bool {
	rc := r.getRedisConn()
	defer rc.Close()

	flag, err := redis.Bool(rc.Do("HEXISTS", hash, key))
	if err != nil {
		return false
	}
	return flag
}


func (r* RRedis) StringSet(key string, value string) {
	rc := r.getRedisConn()
	defer rc.Close()

	_, err := rc.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis list failed:", err)
	}
}

func (r* RRedis) StringGet(key string) (string, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	value, err := redis.String(rc.Do("GET", key))
	if err != nil {
		//fmt.Println("redis list failed:", err)
		return "", err
	}
	return value, nil
}

func (r* RRedis) StringExists(key string) bool {
	rc := r.getRedisConn()
	defer rc.Close()

	flag, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		//fmt.Println("redis list failed:", err)
		return false
	}
	return flag
}

func (r* RRedis) IsNill(err interface{}) bool {
	if err == redis.ErrNil {
		return true
	} else {
		return false
	}
}