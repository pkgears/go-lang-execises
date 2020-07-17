package main // package name
// package any name

// Imports
import "fmt"

// main function
func main() {
	var message string = "Hello, world"
	easyMessage := "Hello, world using :="

	fmt.Println(message)
	fmt.Println(easyMessage)

	// Data types
	// Floats
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)
	// Integers
	var c int = 10
	d := 3
	fmt.Println(c / d)
	// Boolean
	x := true
	y := false
	fmt.Println(x || y)
	fmt.Println(y && x)
	fmt.Println(!x)

	privateFunc()
	PublicFunc()
	printHello()
}

func privateFunc() {
	fmt.Println("For execute this function is necessary export it")
}

// PublicFunc print a message
func PublicFunc() {
	fmt.Println("This function is public")
}

func printHello() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
