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

const (
	RespTypeOk    = "ok"
	RespTypeError = "error"
	VTypeString   = "string"
	VTypeBinary   = "binary"
)

type Resp struct {
	Type    string
	Content string
}

func (r *Resp) Format() string {
	if r.Type == RespTypeOk {
		return fmt.Sprintf("%s", r.Content)
	} else {
		return fmt.Sprintf("%s, %s", r.Type, r.Content)
	}
}

func (r *Resp) FromItem(item *memcache.Item, t string) {
	res := ""
	switch parseVType(t) {
	case VTypeBinary:
		res = fmt.Sprintf("%v", item.Value)
	case VTypeString:
		res = string(item.Value)
	}
	r.Type = RespTypeOk
	r.Content = fmt.Sprintf("%s [flags: %d]", res, item.Flags)
}

func parseVType(t string) string {
	switch t {
	case "binary", "bin":
		return VTypeBinary
	default:
		return VTypeString
	}
}

func (r *Resp) FromErr(err error) {
	if err == memcache.ErrCacheMiss {
		r.Type = RespTypeOk
		r.Content = "(nil)"
		return
	}
	r.Type = RespTypeError
	r.Content = fmt.Sprintf("%v", err)
	return
}

func (d *Dao) GetWithType(key, t string) (resp *Resp) {
	resp = new(Resp)
	item, err := d.Get(key)
	if err != nil {
		resp.FromErr(err)
		return
	}
	resp.FromItem(item, t)
	return
}

func (d *Dao) Get(key string) (item *memcache.Item, err error) {
	item, err = d.mc.Get(key)
	return
}

func (d *Dao) Set(key string, value []byte, exp int32, flags uint32) (err error) {
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: exp,
		Flags:      flags,
	}
	err = d.mc.Set(item)
	return
}

func (d *Dao) Delete(key string) (err error) {
	err = d.mc.Delete(key)
	return
}
