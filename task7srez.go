package main

import "fmt"

func main() {
	var a []int

	for i := 0; i < 25; i++ {
		a = append(a, i*10)
	}
	for k, a := range a {
		println(k, a)
	}
	for _, i := range a {
		fmt.Println(i, a)
	}

	slace := []string{"Potato", "Tomato", "Capychino", "Senirioto", "Bombalelo"}

	for i, elem := range slace {
		fmt.Printf("Index %d Name: %s\n", i, elem)
	}

	slice := slace[1:3]
	for _, i := range slice {
		fmt.Println(i, a)
	}
}
