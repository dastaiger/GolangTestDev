package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		name: "host",
		age:  12,
	}
	saySomething(&p1)
}

func (p *person) speak() {
	fmt.Println(p.age, p.name)
}

func saySomething(h human) {
	h.speak()
}
