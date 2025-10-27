package main

import "fmt"

type avto struct {
	name    string
	surname string
}

func main() {
	avto := make(map[string]int)
	avto["Nissan"] = 5
	avto["Lada"] = 4
	avto["Chevrolet"] = 3
	avto["Honda"] = 2
	avto["Gelly"] = 1

	fmt.Println("Ключ и значение")
	for k, avtos := range avto {
		fmt.Printf("%s, %d\n", k, avtos)
	}

	avtos := map[string]int{
		"Nissan":    1,
		"Lada":      2,
		"Chevrolet": 3,
		"Honda":     4,
		"Geely":     5,
	}

	fmt.Println("Ключ и значенние")
	for k, avto := range avtos {
		fmt.Printf("%s -> %d\n", k, avto)
	}
}
