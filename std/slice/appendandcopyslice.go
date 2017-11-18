package main

import (
	"echo.com/go/utils"
)

func main() {
	var numbers []int
	utils.PrintSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	utils.PrintSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	utils.PrintSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	utils.PrintSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers)) * 2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	utils.PrintSlice(numbers1)
}
