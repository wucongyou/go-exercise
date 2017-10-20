package main

import (
	"flag"
	"fmt"

	"go-exercise/memcache/dao"
)

func main() {
	var (
		host  string
		port  int64
		vtype string
	)
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.Int64Var(&port, "p", 11211, "port")
	flag.StringVar(&vtype, "t", "string", "value type")
	flag.Parse()
	fmt.Printf("connecting to %s:%d\n", host, port)
	d := dao.New(&dao.Conf{
		Host: host,
		Port: port,
	})
	d.Set("ping", []byte("pong"), 1800)
	res, err := d.Get("ping")
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %s\n", res)
	strRes, err := d.FormatWithType("ping", "string")
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %s", strRes)
}
