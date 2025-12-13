package main

import (
	"errors"
	"fmt"
)

func main() {
	var result, reminder, err = intDivision(10, 1)

	if err != nil {
		fmt.Println((err.Error()))
	} else {
		fmt.Printf("Result: %v, Reminder: %v\n", result, reminder)
	}

}

func intDivision(a, b int) (int, int, error) {
	var err error

	if b == 0 {
		err = errors.New("division by zero is not allowed")
		return 0, 0, err
	}

	return a / b, a % b, err
}
