package main

import "fmt"

func main() {
	var a int
	var b string
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var (
		aa = 5
		bb = "Строка"
		cc = true
	)

	fmt.Println(aa, bb, cc)

	var p float32 = 3.1415
	fmt.Println("Pi", p)
}
