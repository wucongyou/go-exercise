package dao

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	_string = "string"
	_byte   = "byte"
)

type Dao struct {
	mc *memcache.Client
}

type Conf struct {
	Host string
	Port int64
}

func New(c *Conf) *Dao {
	return &Dao{
		mc: memcache.New(fmt.Sprintf("%s:%d", c.Host, c.Port)),
	}
}

func (d *Dao) FormatWithType(key, t string) (res string, err error) {
	data, err := d.Get(key)
	switch t {
	case _string:
		res = string(data)
	case _byte:
		res = fmt.Sprintf("%v", data)
	}
	return
}

func (d *Dao) Get(key string) (res []byte, err error) {
	item, err := d.mc.Get(key)
	if err != nil {
		return
	}
	res = item.Value
	return
}

func (d *Dao) Set(key string, value []byte, exp int32) (err error) {
	d.mc.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: exp,
	})
	return
}

func (d *Dao) Delete(key string) (err error) {
	err = d.mc.Delete(key)
	return
}
