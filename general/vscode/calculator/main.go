package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("number 1: ")
	n1, _ := reader.ReadString('\n')       // 10\n
	n1 = strings.Replace(n1, "\n", "", -1) // 10
	number1, _ := strconv.Atoi(n1)

	fmt.Printf("First number: %v\n", number1)

	fmt.Print("number 2: ")
	n2, _ := reader.ReadString('\n')
	n2 = strings.Replace(n2, "\n", "", -1)
	number2, _ := strconv.Atoi(n2)

	fmt.Printf("Second number: %d\n", number2)

	fmt.Print("operator: ")
	operator, _ := reader.ReadString('\n')             // +\n
	operator = strings.Replace(operator, "\n", "", -1) // +
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
