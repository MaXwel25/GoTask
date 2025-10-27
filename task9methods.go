package main

import "fmt"

type Baseball struct {
	name   string
	hits   int
	atBats int
}

func (a *Baseball) average() float64 {
	if a.atBats == 0 {
		return 0
	}
	return float64(a.hits) / float64(a.atBats)
}

func main() {
	users := []Baseball{
		{name: "Maxim", atBats: 100, hits: 20},
		{name: "Aleksey", atBats: 120, hits: 35},
		{name: "Gadzi", atBats: 90, hits: 15},
		{name: "Nikita", atBats: 110, hits: 22},
		{name: "Rodion", atBats: 130, hits: 37},
	}

	for _, user := range users {
		average_sr := user.average()
		fmt.Printf("Игрок: %s\n", user.name)
		fmt.Printf(" Ударов: %d, попаданий: %d\n", user.atBats, user.hits)
		fmt.Printf(" Средний показатель: %.3f\n\n", average_sr)
	}
}
