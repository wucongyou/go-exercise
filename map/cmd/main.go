package main

import "fmt"

func main() {
	m := make(map[string]int, 0)
	fmt.Printf("pm: %p\n", m)
	n := make(map[string]int, 0)
	fmt.Printf("pn: %p\n", n)
	m = n
	fmt.Printf("pm: %p\n", m)

	fmt.Printf("ppm: %p\n", &m)
	fmt.Printf("ppn: %p\n", &n)
	fmt.Printf("ppm: %p\n", &m)
}
