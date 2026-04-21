package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetNumber(prompt string, reader *bufio.Reader, operation ...string) float64 {
	for {
		fmt.Print(prompt)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		if len(operation) > 0 && operation[0] == "/" && number == 0 {
			fmt.Println("Invalid input. Division by zero")
			continue
		}
		return number
	}
}

func GetOperation(prompt string, reader *bufio.Reader) string {
	for {
		fmt.Print(prompt)

		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		switch op {
		case "+", "-", "*", "/":
			return op
		default:
			fmt.Println("Invalid input")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	lhs := GetNumber("Input left operand: ", reader)
	operation := GetOperation("Input operation (+, -, *, /): ", reader)
	rhs := GetNumber("Input right operand: ", reader, operation)

	var result float64

	switch operation {
	case "+":
		result = lhs + rhs
	case "-":
		result = lhs - rhs
	case "*":
		result = lhs * rhs
	case "/":
		result = lhs / rhs
	}

	fmt.Printf("Result : %.03f", result)
}
