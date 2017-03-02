package main

import "fmt"

var (
	zids []int64 = make([]int64, 0)
)

type Person struct {
}

func getPerson() (p Person) {
	return
}

func getPerson2() Person {
	var p Person
	return p
}

func main() {
	fmt.Printf("addr: %p,len: %d,cap: %d\n", &zids, len(zids), cap(zids))
	zids = append(zids, 1)
	fmt.Printf("addr: %p,len: %d,cap: %d\n", &zids, len(zids), cap(zids))
	fmt.Println(getPerson())
	fmt.Println(getPerson2())
}
