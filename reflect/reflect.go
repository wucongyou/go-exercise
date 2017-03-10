package reflect

import (
	"fmt"
)

type SayHello interface {
	SayHello()
}

type Person struct {
}

func (f *Person) SayHello() {
	fmt.Println("hello from person")
}

type Proxy struct {
	Person
}

func (f *Proxy) SayHello() {
	fmt.Println("hello from proxy")
}

func Client() {
	var (
		p = new(SayHello)
	)
	*p = new(Person)
	(*p).SayHello()
	*p = new(Proxy)
	(*p).SayHello()
}
