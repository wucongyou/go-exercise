package namedret

import "fmt"

type Person struct {
}

func Client() {
	fmt.Println(getPerson())
	fmt.Println(getPerson2())
}

func getPerson() (p Person) {
	return
}

func getPerson2() Person {
	var p Person
	return p
}
