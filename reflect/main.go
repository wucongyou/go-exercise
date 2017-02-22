package main

import (
	"fmt"
	"reflect"
)

type Person struct {
}

func (f *Person) SayHello() string {
	return "hello"
}

func when(m reflect.Method) {
	return
}

func thenReturn(v interface{}) {

}

func mock(a interface{}) (res interface{}) {
	for m := 0; m < reflect.TypeOf(a).NumMethod(); m++ {
		method := reflect.TypeOf(a).Method(m)
		fmt.Println(method.Type)         // func(*main.MyStruct) string
		fmt.Println(method.Name)         // GetName
		fmt.Println(method.Type.NumIn()) // 参数个数
		fmt.Println(method.Type.In(0))   // 参数类型
	}
	return
}

func main() {
	var (
		t reflect.Type
		v reflect.Value
	)
	s := "hello"
	t = reflect.TypeOf(s)
	fmt.Printf("t: %v\n", t)
	v = reflect.ValueOf(s)
	fmt.Printf("v: %v\n", v)
	p := new(Person)
	t = reflect.TypeOf(p)
	fmt.Printf("t: %v\n", t)
	v = reflect.ValueOf(p)
	fmt.Printf("v: %v\n", v)
	mock(new(Person))
}
