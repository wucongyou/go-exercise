package main

import (
	"flag"
	"fmt"
)

var ip string

func init()  {
	flag.StringVar(&ip, "ip", "noip", "ip address")
}

func main() {
	flag.Parse()
	fmt.Printf("ip: %s", ip)
}
