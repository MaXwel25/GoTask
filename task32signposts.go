package main

import "fmt"

type Info struct {
	a int
	b string
	c float32
}

func change_a(i *Info, newa int) {
	i.a = newa
}

func change_b(i *Info, newb string) {
	i.b = newb
}

func change_c(i *Info, newc float32) {
	i.c = newc
}

func main() {
	info := Info{
		a: 5,
		b: "Автомобиль",
		c: 1.5,
	}

	fmt.Printf("\na = %d, b = %s, c = %s\n", info.a, info.b, info.c)

	change_a(&info, 100)
	change_b(&info, "Мотоцикл")
	change_c(&info, 0.8)
	fmt.Printf("\na = %d, b = %s, c = %s\n", info.a, info.b, info.c)
}
