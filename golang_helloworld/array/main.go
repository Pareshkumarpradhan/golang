package main

import "fmt"

func main() {
	fmt.Println("we are learning array in GoLang")

	var name[5]string

	name[0] = "Rahul"
	name[1] = "Rohit"

	fmt.Println("Name is: ", name)

	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is: ", number)

	fmt.Println("Length of the array is: ", len(number))

	fmt.Println("First element of the array is: ", number[0])
}