package dao

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func New(c *Conf) *Dao {
	return &Dao{
		mc: memcache.New(fmt.Sprintf("%s:%d", c.Host, c.Port)),
	}
}

type Dao struct {
	mc *memcache.Client
}

type Conf struct {
	Host string
	Port int64
}

func (d *Dao) Get(key string) (*memcache.Item, error) {
	return d.mc.Get(key)
}

func (d *Dao) Set(item *memcache.Item) error {
	return d.mc.Set(item)
}

func (d *Dao) Delete(key string) error {
	return d.mc.Delete(key)
}
