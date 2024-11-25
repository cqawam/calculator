package calculator

import "fmt"


func Add(num1 float64, num2 float64) (sum float64) {
	return num1 + num2
}

func Subtract(num1 float64, num2 float64) (diff float64) {
	return num1 - num2
}

func Multiply(num1 float64, num2 float64) (product float64) {
	return num1 * num2
}

func Divide(num1 float64, num2 float64) (quotient float64, err error) {
	if num2 != 0 {
		return num1 / num2, nil
	}
	return 0, fmt.Errorf("Division by zero")
}