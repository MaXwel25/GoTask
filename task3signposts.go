package main

import "fmt"

func main() {
	var a int = 20

	fmt.Printf("Адрес переменной: %p\n", &a)
	fmt.Printf("Значение переменной: %d\n", a)
}
