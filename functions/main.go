package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func multiplication(a, b int) int {
	return a * b
}
func main() {
	fmt.Println(greeting("Semre"))
	fmt.Println(multiplication(3, 4))
}
