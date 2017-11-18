package main

import (
	"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 	`error` 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	numerator: %d
	denominator: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(numerator int, denominator int) (result int, errorMsg string) {
	if denominator == 0 {
		dData := DivideError{
			dividee: numerator,
			divider: denominator,
		}
		errorMsg = dData.Error()
		return
	} else {
		return numerator / denominator, ""
	}

}

func main() {

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
