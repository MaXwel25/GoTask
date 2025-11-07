package main

import "fmt"

func main() {
	tt := []string{"", "G", "bob", "otto", "汉字汉", "test"}

	for _, input := range tt {
		success := IsPalindrome(input)

		switch success {
		case true:
			fmt.Printf("%q is a palindrome\n", input)

		case false:
			fmt.Printf("%q is NOT a palindrome\n", input)
		}
	}
}

func IsPalindrome(input string) bool {

	if input == "" || len(input) == 1 {
		return true
	}

	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
