package main

import "fmt"

func main() {
	// Input two numbers and an operator from user
	var num1, num2 int
	var operator string

	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&num1, &num2)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Switch statement to calculate the result of two numbers
	// based on chosen operator
	switch operator {
	case "+":
		fmt.Printf("%d + %d = %d", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d", num1, num2, num1/num2)
	default:
		fmt.Println("Invalid operator")
	}
}
