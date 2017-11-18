package main

import (
	"echo.com/go/utils"
)

func main() {

	var a uint = 60        /* 60 = 0011 1100 */
	var b uint = 13        /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b       /* 12 = 0000 1100 */
	utils.PrintResultUint(a, b, "&", c)

	c = a | b       /* 61 = 0011 1101 */
	utils.PrintResultUint(a, b, "|", c)

	c = a ^ b       /* 49 = 0011 0001 */
	utils.PrintResultUint(a, b, "^", c)

	c = a << 2     /* 240 = 1111 0000 */
	utils.PrintResultUint(a, 2, "<<", c)

	c = a >> 2     /* 15 = 0000 1111 */
	utils.PrintResultUint(a, 2, ">>", c)
}
