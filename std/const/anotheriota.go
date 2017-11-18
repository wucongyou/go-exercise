package main

const (
	i = 1 << iota
	j = 3 << iota
	k // 3 << 2
	l // 3 << 3
)

func main() {
	println(i, j, k, l)
}
