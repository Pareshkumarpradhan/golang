package main

import "fmt"

func simpleFunction(){
	fmt.Println("This is a simple function")
}

func add(a,b int) int {
	return a+b
}

func main() {
	fmt.Println("we are learning functions in golang")
	simpleFunction()

	ans := add(2,3)
	fmt.Println("The sum is: ", ans)
}