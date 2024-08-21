package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)

	// // define map
	// emails["semre"] = "semre@gmail.com"
	// emails["dave"] = "dave@gmail.com"

	// declare and define map
	// emails := map[string]string{"semre": "semre@gmail.com", "dave": "dave@gmail.com"}

	// fmt.Println(emails)

	ids := []int{23, 25, 27, 29, 31, 33}

	for i, id := range ids {
		fmt.Printf("%d-ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

}
