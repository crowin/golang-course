package cache

import (
	"time"
)

type Cache struct {
	data map[string]Value
}

type Value struct {
	value      any
	expiration time.Time
}

func (c *Cache) Set(key string, value any) {
	c.data[key] = Value{
		value:      value,
		expiration: time.Now().Add(time.Hour * time.Duration(24)),
	}
}

func (c *Cache) Get(key string) (value any, found bool) {
	result, exist := c.data[key]
	if !exist {
		return value, false
	} else if result.expiration.Before(time.Now()) {
		delete(c.data, key)
		return value, false
	} else {
		return result.value, true
	}
}

func (c *Cache) Remove(key string) {
	delete(c.data, key)
}

func NewCache() *Cache {
	c := &Cache{
		data: make(map[string]Value),
	}

	return c
}
