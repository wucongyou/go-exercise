package main

import "fmt"

type Dao struct {
	Url string
}

func (d *Dao) url() {
	fmt.Println(d.Url)
}
