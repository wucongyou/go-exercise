package main

import (
	"fmt"
	"strings"
	"os"
)

func exe1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func exe2() {
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		s += sep + fmt.Sprint(index) + " " + arg
		sep = "\n"
	}
	fmt.Println(s)
}

func main() {
	exe1()
	exe2()
}
