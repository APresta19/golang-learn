package main

import (
	"fmt"
	"errors"
)

func divideNums(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return a / b, nil
}

func main() {
	calc, err := divideNums(5, 2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} 
	fmt.Println(calc)
}