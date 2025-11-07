package main

import "fmt"

func main() {
	tt := []int{-1, 1, 9, 10, 1001, 125}

	for _, input := range tt {
		success := IsPalindrome(input)

		switch success {
		case true:
			fmt.Printf("%d is a palindrome\n", input)

		case false:
			fmt.Printf("%d is NOT a palindrome\n", input)
		}
	}
}

func IsPalindrome(input int) bool {

	if input < 0 {
		return false
	}
	if input >= 0 && input < 10 {
		return true
	}

	rev := Reverse(input)

	return input == rev
}

func Reverse(num int) int {

	var result int

	for num != 0 {

		last := num % 10

		result = result * 10

		result += last
		num = num / 10
	}

	return result
}
