package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	pool := newPool("127.0.0.1:6379")
	conn := pool.Get()

	// resp, err := redis.Int(conn.Do("GET", "test"))
	resp, err := conn.Do("SET", "test", "1")
	resp1, err := redis.String(resp, err)
	fmt.Println(err)
	fmt.Println(resp1)
	conn.Close()

	p1 := newTestPool("127.0.0.1:6379")

	conn1, err := p1.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn1.Close()
	resp2, err := redis.Int(conn1.Do("GET", "test"))
	fmt.Println(resp2)
	fmt.Println(err)

}

type Demo struct {
	Name string
	f    func()
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func newTestPool(addr string) *pool {
	return &pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}
