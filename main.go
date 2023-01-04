package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	s := NewStore(NewRedisCache(client, time.Second*4))

	for q := 0; q < 2; q++ {
		val, err := s.Get(1)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(val)
		time.Sleep(5 * time.Second)
	}

}
