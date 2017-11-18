package main

import (
	"echo.com/go/utils"
)

func main() {
	var a int = 21
	var b int = 10
	var c bool

	c = a == b
	utils.PrintResultBool(a, b, "==", c)
	c = a != b
	utils.PrintResultBool(a, b, "!=", c)
	c = a < b
	utils.PrintResultBool(a, b, "<", c)
	c = a >= b
	utils.PrintResultBool(a, b, ">=", c)
	c = a > b
	utils.PrintResultBool(a, b, ">", c)
	c = a <= b
	utils.PrintResultBool(a, b, "<=", c)

	/* Lets change value of a and b */
	a = 5
	b = 20
	c = a <= b
	utils.PrintResultBool(a, b, "<=", c)
	c = b >= a
	utils.PrintResultBool(b, a, ">=", c)
}
