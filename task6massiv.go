package main

func main() {
	var fruits [5]string
	info := [5]string{"Apple", "Orange", "Banana", "Grape", "Plum"}

	fruits = info
	for i, fruit := range fruits {
		println(fruit, &fruits[i])
	}

	for i, fruit := range fruits {
		println(i, fruit)
	}
}
