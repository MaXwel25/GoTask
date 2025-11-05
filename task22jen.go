package main

import (
	"encoding/json"
	"fmt"
)

func marshall[T json.Marshaler](value T) ([]byte, error) {
	return value.MarshalJSON()
}

type user struct {
	Name string
	Email string
}

func (u user) MarshalJSON() ([]byte, error) {
	type userJSON struct {
		Name string 'json:"name"'
		Email string 'json:"email"'
	}

	custom := userJSON{
		Name: u.Name,
		Email: u.Email,
	}
	return json.Marshal(custom)
}

func main() {
	u := user {
		Name: "Alice",
		Email: "alice@mail.ru"
	}

	jsonData, err := marshal(u)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("JSON: %s\n", string(jsonData))
}
