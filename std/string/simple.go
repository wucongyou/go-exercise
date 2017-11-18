package main

import "fmt"

func main() {
	str := "这是utf-8编码"
	r := []rune(str)
	for _, item := range r {
		fmt.Println(item, string(item))
	}
}
