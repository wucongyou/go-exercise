package model

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

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
		return fmt.Sprint(r.Content)
	} else {
		return fmt.Sprintf("%s, %s", r.Type, r.Content)
	}
}

func FromItem(item *memcache.Item, t string) (r *Resp) {
	switch parseVType(t) {
	case VTypeBinary:
		return &Resp{
			Type:    RespTypeOk,
			Content: fmt.Sprintf("%v [flags: %d]", item.Value, item.Flags),
		}
	case VTypeString:
		return &Resp{
			Type:    RespTypeOk,
			Content: fmt.Sprintf("%s [flags: %d]", string(item.Value), item.Flags),
		}
	}
	return
}

func parseVType(t string) string {
	switch t {
	case "binary", "bin":
		return VTypeBinary
	default:
		return VTypeString
	}
}

func FromErr(err error) (r *Resp) {
	if err == memcache.ErrCacheMiss {
		return &Resp{
			Type:    RespTypeOk,
			Content: "(nil)",
		}
	}
	return &Resp{
		Type:    RespTypeError,
		Content: fmt.Sprint(err),
	}
}
