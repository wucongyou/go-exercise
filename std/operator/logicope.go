package main

import (
	"echo.com/go/utils"
)

func main() {
	var a bool = true
	var b bool = false
	var c bool
	c = a && b
	utils.PrintResultLogic(a, b, "&&", c)
	c = a || b
	utils.PrintResultLogic(a, b, "||", c)
	c = b || a
	utils.PrintResultLogic(b, a, "||", c)
	/* 修改 a 和 b 的值 */
	a = true
	b = true
	c = a && b
	utils.PrintResultLogic(a, b, "&&", c)
}
