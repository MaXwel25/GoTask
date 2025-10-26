package main

import (
	"errors"
	"fmt"
)

type User struct {
	Age   int
	Name  string
	Phone string
}

func create_user(age int, name string, phone string) (*User, error) {
	if age > 100 || age < 0 {
		return nil, errors.New("Чушь собачья)")
	}
	if len(phone) != 11 {
		return nil, errors.New("хз какой номер)")
	}

	user := &User{
		Age:   age,
		Name:  name,
		Phone: phone,
	}
	return user, nil
}

func main() {
	user, err := create_user(20, "Иван", "88005553535")
	if err != nil {
		fmt.Printf("Ошибка при создании пользователя: %v\n", err)
		return
	}

	fmt.Printf("Польватель создан:\n")
	fmt.Printf(" Возраст: %d\n", user.Age)
	fmt.Printf(" Имя: %s\n", user.Name)
	fmt.Printf(" Телефон: %s\n", user.Phone)

	_, err = create_user(25, "", "880053515")
	if err != nil {
		fmt.Printf(" Ошибка: %v\n", err)
	} else {
		fmt.Println(" Пользователь создан")
	}
}
