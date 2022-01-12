package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// func redis_conn() {
// 	setPasswd := redis.DialPassword("") //设置密码
// 	// tcp连接redis
// 	conn, err := redis.Dial("tcp", "", setPasswd)
// 	if err != nil {
// 		fmt.Println("connect redis error:", err)
// 		return
// 	}
// }

func main() {
	setPasswd := redis.DialPassword("111111") //设置密码
	conn, err := redis.Dial("tcp", "47.114.171.118:32336", setPasswd)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok")
	}

	defer conn.Close()

	// _, err = conn.Do("ZADD", "mykey", "INCR", 1, "robot1")
	set := "mykey"
	key := "robot4"
	value := 35
	// _, err = conn.Do("ZADD", set, value, key)
	// if err != nil {
	// 	fmt.Println("redis set failed:", err)
	// }

	keys, err := redis.Strings(conn.Do("ZADD", set, value, key))
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	fmt.Println("keys : ", keys)

	// _, err = conn.Do("ZADD", "mykey", "INCR", 1, "robot2")
	// if err != nil {
	// 	fmt.Println("redis set failed:", err)
	// }

	user_map, err := redis.StringMap(conn.Do("ZRANGE", "mykey", 0, 10, "withscores"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get myKey: %v \n", user_map)
	}

	for user := range user_map {
		fmt.Printf("user name: %v %v\n", user, user_map[user])
	}
}
