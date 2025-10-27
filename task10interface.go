package main

import "fmt"

type speaker interface {
	speak()
}

type english struct {
	name string
}

func (e english) speak() {
	fmt.Println("Hello World")
}

type chinese struct {
	name string
}

func (c chinese) speak() {
	fmt.Println("sdfewfweсцвсы")
}

func sayHello(sp speaker) {
	fmt.Print("sayHello: ")
	sp.speak()
}

func main() {
	var sp speaker
	sp = &english{name: "Maxim"}
	sp.speak()

	sp = &chinese{name: "dsdсвсц!№"}
	sp.speak()

	eng := english{name: "Maria"}
	chi := chinese{name: "1d3сыц"}

	sayHello(&eng)
	sayHello(&chi)
}
