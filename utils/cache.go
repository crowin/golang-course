package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data  map[string]Value
	mutex sync.RWMutex
	quit  chan struct{}
}

type Value struct {
	value      any
	expiration int64
}

func (c *Cache) Set(key string, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = Value{
		value:      value,
		expiration: time.Now().Add(time.Second * time.Duration(15)).UnixMilli(),
	}
	fmt.Println("Value was added to cache")
}

func (c *Cache) Get(key string) (value any, isExist bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result, exist := c.data[key]
	if !exist {
		return nil, false
	} else if time.Now().UnixMilli() > result.expiration {
		c.remove(key)
		return nil, false
	}

	fmt.Println("Value was gotten from cache")
	return result.value, true
}

func (c *Cache) remove(key string) {
	fmt.Printf("Value with key %s will be deleted\n", key)
	delete(c.data, key)
}

func (c *Cache) cleanUp() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			fmt.Printf("Cleaning is starting\n")

			for k, v := range c.data {
				if time.Now().UnixMilli() > v.expiration {
					c.remove(k)
				}
			}

			c.mutex.Unlock()
		case <-c.quit:
			return
		}
	}
}

func (c *Cache) StopCleanUp() {
	close(c.quit)
	fmt.Println("Stop cleaning")
}

func NewCache() *Cache {
	c := &Cache{
		data: make(map[string]Value),
		quit: make(chan struct{}),
	}

	go c.cleanUp()

	return c
}
