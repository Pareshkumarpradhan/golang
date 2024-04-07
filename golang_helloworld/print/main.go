package main

import "fmt"

func main() {
	age := 24
	name := "Paresh"
	height := 5.7

	fmt.Println("age: ", age, "name: ", name, "height: ", height)
	fmt.Println("hello world")

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)

	fmt.Printf("type og age is %T\n", age)
}