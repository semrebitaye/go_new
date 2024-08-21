package main

import "fmt"

func main() {
	fmt..Println(devide(3, 0))
}

func devide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("can not devide by zero")
	}
	return a / b, nil
}
