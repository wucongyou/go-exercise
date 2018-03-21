package main

import "fmt"

func main() {
	const imax  =int64(0xffffffffffffffff)
	now := int64(1514297068)
	a := int64(0x1ffffffff) // int64_max - ts
	b := uint32(a)          // will be 0xffffffff
	fmt.Printf("now: %x, a: %x = %d, b: %x = %d", now, a, a, b, b)
}
