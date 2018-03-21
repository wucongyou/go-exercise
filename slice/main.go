package main

import "fmt"

func main() {
	var s []int
	s[0] = 1
	s = append(s, 1)
	fmt.Print(s)
}
