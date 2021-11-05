package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func (r *RRedis) ZSetAdd_test() []string {

	rc := r.getRedisConn()

	fmt.Println(rc)
	defer rc.Close()

	keys, err := redis.Strings(rc.Do("KEYS", "*"))
	if err != nil {
		return make([]string, 0)
	}
	return keys
}
