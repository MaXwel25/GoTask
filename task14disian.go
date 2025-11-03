package main

import (
	"fmt"
	"errors"
)

var (
	ErrInvalidValue = errors.New("invalid value: amount cannot be Zero")
	ErrAmountTooLarge = errors.New("amount too large: cannot exceed $1000")
)

func checkAmount(amount float64) error {
	if amount ==0 {
		return ErrInvalidValue
	}
	if amount > 1000 {
		return ErrAmountTooLarge
	}

	return nil
}

type appError struct {
	err error
	message string
	code int
}

func (a *appError) Error() string {
	return fmt.Sprintf("code %d: %s - %v", a.code, a.message, a.err)
}

func (a *appError) Temporary() bool {
	return a.code != 9
}

type temporary interface {
	Temporary() bool
}

func checkFlag(flag bool) error {
	if !flag {
		return &appError{
			err: errors.New("validation failed"),
			message: "invalid input data",
			code: 9,
		}
	}
	return errors.New("default error occurred")
}

func main() {
	amounts := []float64{0, 1500, 500}

	for _, amt := range amounts {
		fmt.Printf("Checking amount: $%.2f\n", amt)

		if err := checkAmount(amt); err != nil {
			switch err {
			case ErrInvalidValue:
				fmt.Println("Error:", err)
			case ErrAmountTooLarge:
				fmt.Println("Error:", err)
			default:
				fmt.Println("Unkown Error:", err)
			}
			continue
		}
		fmt.Println("Amount is valid")
	}

	fmt.Println("Testing with flag = false:")
	err := checkFlag(false)
	if err != nil {
		fmt.Printf("Errors: %v\n", err)

		if tempErr, ok := err.(temporary); ok {
			fmt.Printf("Is temporary: %t\n", tempErr.Temporary())
		} else {
			fmt.Println("Errors does not implement temporary interface")
		}
	}

	fmt.Println("\nTesting with flag = true:")
	err = checkFlag(true)
	if err != nil {
		fmt.Printf("Errors: %v\n", err)

		if tempErr, ok := err.(temporary); ok {
			fmt.Printf("Is temporary: %t\n", tempErr.Temporary())
		} else {
			fmt.Println("Error does not implement temporary interface")
		}
	}

}
