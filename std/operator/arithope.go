package main

import (
	"fmt"
	"echo.com/go/utils"
)

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	utils.PrintResultInt("+", c)
	c = a - b
	utils.PrintResultInt("-", c)
	c = a * b
	utils.PrintResultInt("*", c)
	c = a / b
	utils.PrintResultInt("/", c)
	c = a % b
	utils.PrintResultInt("%", c)
	a++
	fmt.Printf("a++ = %d\n", a)
	a--
	fmt.Printf("a-- = %d\n", a)
}
