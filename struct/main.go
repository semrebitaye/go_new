package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Sex  string
	Age  int
}

func (p *Person) great() string {
	return "Hello, my name is " + p.Name + " and i am " + strconv.Itoa(p.Age) + " years old"
}

func (p *Person) hasBirthday() {
	p.Age++
}

func main() {
	person1 := Person{Name: "semre", Sex: "Male", Age: 28}
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.great())

}
