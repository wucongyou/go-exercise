package main

// 全局变量允许赋值当时不使用；
var (
	globalFoo = "foo"
	globalBar = true
)

func main() {
	// 局部变量赋值之后必须使用
	// foo := 7
	// _ 用于抛弃值
	_, bar := 5, 7

	// 值传递与引用传递
	// 基本类型属于值类型：int ,float,bool,string

	// 允许单行多变量声明、赋值
	var a, b, c int
	a, b, c = 1, 2, 3
	println(bar)
	println(a, b, c)
}
