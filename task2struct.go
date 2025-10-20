package main

import "fmt"

func main() {
	type easy struct {
		flag    bool
		counter int16
		pi      float32
	}

	var e1 easy
	fmt.Printf("%+v\n", e1)

	type user struct {
		name  string
		email string
		age   int
	}

	max := user{
		name:  "Maxim",
		email: "maxim@mail.ru",
		age:   20,
	}

	fmt.Println("Имя: ", max.name)
	fmt.Println("Почта: ", max.email)
	fmt.Println("Возраст: ", max.age)

	anonim := struct {
		name  string
		email string
		age   int
	}{
		name:  "Gadzi",
		email: "gadzi05@gmail.com",
		age:   19,
	}

	fmt.Println("Имя ", anonim.name)
	fmt.Println("Почта ", anonim.email)
	fmt.Println("Возраст ", anonim.age)
}
