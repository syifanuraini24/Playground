package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var operator string
	var number1, number2 int

	n1 := os.Getenv("FIRST_NUMBER")
	number1, _ = strconv.Atoi(n1)

	fmt.Printf("First number: %v\n", number1)

	n2 := os.Getenv("SECOND_NUMBER")
	number2, _ = strconv.Atoi(n2)

	fmt.Printf("Second number: %d\n", number2)

	operator = os.Getenv("ACTION")
	fmt.Printf("Selected Operator (+,-,/,%%,*): %s\n", operator)

	output := 0
	switch operator {
	case "+":
		output = number1 + number2
	case "-":
		output = number1 - number2
	case "*":
		output = number1 * number2
	case "/":
		output = number1 / number2
	case "%":
		output = number1 % number2
	default:
		fmt.Println("Invalid Operation")
	}
	fmt.Printf("%d %s %d = %d", number1, operator, number2, output)
}
