package main

import "fmt"

func modifyValueByReferance(num *int) {
	*num = *num + 10
}

func main() {
	// var num int
	num := 2

	// var ptr *int
	// ptr = &num

	ptr := &num

	// fmt.Println("Num has value: ", num)
	fmt.Println("pointer contains: ", ptr)
	fmt.Println("Data contains through Pointer: ", *ptr)

	var pointer *int  // default pointer == nil
	if pointer == nil {
          fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyValueByReferance(&value)
	fmt.Println("Value contains :", value)
}
