package main

import (
	"fmt"
	"reflect"
)

type SayHello interface {
	SayHello() string
}

type Person struct {
}

func (f *Person) SayHello() string {
	fmt.Println("hello")
	return "hello"
}

type Proxy struct {
	Person
}

func (f *Proxy) SayHello() string {
	fmt.Println("hello from proxy")
	return "hello from proxy"
}

func methods(a interface{}) {
	for i := 0; i < reflect.TypeOf(a).NumMethod(); i++ {
		method := reflect.TypeOf(a).Method(i)
		fmt.Println(method.Type)         // func(*main.MyStruct) string
		fmt.Println(method.Name)         // GetName
		fmt.Println(method.Type.NumIn()) // 参数个数
		fmt.Println(method.Type.In(0))   // 参数类型
	}
}

func main() {
	var (
		t reflect.Type
		v reflect.Value
		p = new(SayHello)
	)
	s := "hello"
	t = reflect.TypeOf(s)
	fmt.Printf("t: %v\n", t)
	v = reflect.ValueOf(s)
	fmt.Printf("v: %v\n", v)
	*p = new(Person)
	t = reflect.TypeOf(p)
	fmt.Printf("t: %v\n", t)
	v = reflect.ValueOf(p)
	fmt.Printf("v: %v\n", v)
	*p = new(Person)
	(*p).SayHello()
	methods(p)
	*p = new(Proxy)
	(*p).SayHello()
}
