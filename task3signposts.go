package main

import "fmt"

func main() {
	var a int = 20

	fmt.Printf("Адрес переменной: %p\n", &a)
	fmt.Printf("Значение переменной: %d\n", a)

	var p *int = &a

	fmt.Printf("Адрес указателя: %p\n", &p)
	fmt.Printf("Значение указателя (адрес a): %p\n", p)
	fmt.Printf("Значение на которое указыввает указатель: %d\n", *p)
}
