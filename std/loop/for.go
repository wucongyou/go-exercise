package main

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	/* 定义局部变量 */
	var i, j int

	// 循环嵌套
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if (i % j == 0) {
				break; // 如果发现因子，则不是素数
			}
		}
		if (j > (i / j)) {
			fmt.Printf("%d  是素数\n", i);
		}
	}

	for true {
		fmt.Printf("这是无限循环。\n");
	}
}
