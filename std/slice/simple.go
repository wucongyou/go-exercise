package main

import (
	"echo.com/go/utils"
	"fmt"
)

func main() {
	var numbers = make([]int, 3, 5)
	utils.PrintSlice(numbers)

	months := [...]string{1: "January", 2:"February", 12: "December"}

	for index, month := range months {
		fmt.Println(index, month)
	}
}
