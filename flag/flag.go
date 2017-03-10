package flag

import (
	"flag"
	"fmt"
)

var ip string

func init() {
	flag.StringVar(&ip, "ip", "noip", "ip address")
}

func Start() {
	flag.Parse()
	fmt.Printf("ip: %s", ip)
}
