package utils

import "fmt"

func PrintResultLogic(a bool, b bool, ope string, result bool) {
	fmt.Printf("%t %s %t = %t\n", a, ope, b, result)
}

func PrintResultBool(a int, b int, ope string, result bool) {
	fmt.Printf("%d %s %d = %t\n", a, ope, b, result)
}

func PrintResultInt(ope string, result int) {
	fmt.Printf("a %s b = %d\n", ope, result)
}

func PrintResultUint(a uint, b uint, ope string, result uint) {
	fmt.Printf("%d %s %d = %d\n", a, ope, b, result)
}

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

