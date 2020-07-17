package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input string, operator string) int {
	clearedInput := strings.Split(input, operator)
	operatorOne := parseInt(clearedInput[0])
	operatorTwo := parseInt(clearedInput[1])
	switch operator {
	case "+":
		fmt.Println(operatorOne + operatorTwo)
		return (operatorOne + operatorTwo)
	case "-":
		fmt.Println(operatorOne - operatorTwo)
		return (operatorOne - operatorTwo)
	case "*":
		fmt.Println(operatorOne * operatorTwo)
		return (operatorOne * operatorTwo)
	case "/":
		fmt.Println(operatorOne / operatorTwo)
		return (operatorOne / operatorTwo)
	default:
		fmt.Println("Unknown operation")
		return 0
	}
}

func parseInt(input string) int {
	value, _ := strconv.Atoi(input)
	return value
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func main() {
	input := readInput()
	operator := readInput()

	c := calc{}
	fmt.Println(c.operate(input, operator))
}
