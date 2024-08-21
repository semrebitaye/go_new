package main

import "fmt"

func main() {
	// // declare array
	// var fruitArr [2]string

	// // assign array
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	// declare and assign at the same time
	// fruitArr := [2]string{"mango", "avocado"}

	// fmt.Println(fruitArr)

	// fruitSlice := []string{"mango", "avocado", "apple"}

	// fmt.Println(len(fruitSlice))

	// x := 8
	// y := 8

	// if x > y {
	// 	fmt.Printf("%d is greater than %d \n", x, y)
	// } else if x == y {
	// 	fmt.Printf("%d is equal to %d \n", x, y)
	// } else {
	// 	fmt.Printf("%d is greater than %d \n", y, x)
	// }

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzBuz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

}
