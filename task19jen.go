package main

import (
	"encoding/json"
	"fmt"
)

func marshal[T any](value T) ([]byte, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type User struct {
	Name string
	Age int
}

func main() {
	user := User{
		Name: "Alice",
		Age: 30,
	}

	data, err := marshal(user)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(string(data))
}
