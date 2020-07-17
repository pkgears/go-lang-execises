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
		return (operatorOne + operatorTwo)
	case "-":
		return (operatorOne - operatorTwo)
	case "*":
		return (operatorOne * operatorTwo)
	case "/":
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
	fmt.Println("Type operation ex(2+2, 4-2, 3*3, 9/3")
	input := readInput()
	fmt.Println("Type operator ex(+ - * /)")
	operator := readInput()

	c := calc{}
	fmt.Println(c.operate(input, operator))
}
