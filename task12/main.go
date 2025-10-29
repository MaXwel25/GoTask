package main

import (
	"fmt"

	"play.ground/toy"
)

func main() {

	t := toy.New("Bat", 28)

	t.UpdateOnHand(100)
	t.UpdateSold(2)

	fmt.Println("Name", t.Name)
	fmt.Println("Weight", t.Weight)
	fmt.Println("OnHand", t.OnHand())
	fmt.Println("Sold", t.Sold())
}
