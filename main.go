package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type Cacher interface {
	Get(int) (string, bool)
	Remove(int) error
	Set(int, string) error
}

type NopCache struct{}

func (c NopCache) Get(int) (string, bool) {
	return "", false
}

func (c NopCache) Set(int, string) error {
	return nil
}

func (c NopCache) Remove(int) error {
	return nil
}

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {

	data := map[int]string{
		1: "Elon musk is the owner of twitter",
		2: "Foo is not bar",
		3: "Must watch show",
	}
	return &Store{
		data:  data,
		cache: c,
	}
}

func (s *Store) getFromCache(key int) (string, bool) {
	val, ok := s.cache.Get(key)

	if !ok {

		fmt.Println("returning key from cache")
		return val, ok
	}
	return val, ok
}

func (s *Store) Get(key int) (string, error) {

	val, err := s.getFromCache(key)
	val, err = s.data[key]

	if !err {
		return "", fmt.Errorf("key not found : %d", key)
	}

	return val, nil
}
func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	s := NewStore(&NopCache{})

	val, err := s.Get(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client, val)

}
