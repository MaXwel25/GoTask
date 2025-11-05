package main

import (
	"errors"
	"fmt"
)

type stack[T any] struct {
	data []T
}

func (s *stack[T]) pop() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]

	s.data = s.data[:lastIndex]

	return value, nil
}

func main() {
	s := stack[int]{}

	s.push(10)
	s.push(20)

	value1, err1 := s.pop()
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Poped value: %d\n", value1)
	}

	value3, err3 := s.pop()

	if err3 != nil {
		fmt.Printf("Error: %v\n", err3)
	} else {
		fmt.Printf("Popped value: %d\n", value3)
	}
}
