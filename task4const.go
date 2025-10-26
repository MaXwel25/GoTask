package main

import "fmt"

func main() {
	const a = 10
	const b int = 25 // типизированная константа
	fmt.Printf("Нетипизированная: ", a)
	fmt.Printf("\nТипизированная: ", b)
	fmt.Printf("\nНетипизированная: %v %T", a)
	fmt.Printf("\nТипизированная: %v %T", b)

	const (
		aa = 1000
		cc = 2
	)

	var res float64 = aa / cc
	fmt.Printf("\nЧислитель: ", aa)
	fmt.Printf("\nЗнаминатель: ", cc)
	fmt.Printf("\nРезультат: ", res)
}
