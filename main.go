package main

import "fmt"

var (
	zids []int64 = make([]int64, 0)
)

func main() {
	fmt.Printf("addr: %p,len: %d,cap: %d\n", &zids, len(zids), cap(zids))
	zids = append(zids, 1)
	fmt.Printf("addr: %p,len: %d,cap: %d\n", &zids, len(zids), cap(zids))
}
