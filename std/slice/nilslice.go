package main

import (
	"fmt"
	"echo.com/go/utils"
)

func main() {
	var numbers []int

	utils.PrintSlice(numbers)

	if (numbers == nil) {
		fmt.Printf("切片是空的")
	}
}
