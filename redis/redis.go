package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

type Item struct {
	Uname  string `json:"uname"`
	Mid    int64  `json:"mid"`
	Type   string `json:"type"`
	Action string `json:"action"`
}

func RedisClient() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	item := &Item{
		Uname:  "John",
		Mid:    1,
		Type:   "update",
		Action: "updatePersonInfo",
	}
	v, err := json.Marshal(item)
	_, err = conn.Do("lpush", fmt.Sprintf("%dGoNotifyQueue", item.Mid%10), v)
	if err != nil {
		fmt.Printf("conn.Do() error(%v)", err)
	}
}
