package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 6, 7, 8, 9, 10)
	// fmt.Println("numbers:", numbers)
	// fmt.Printf("numbers has data type: %T\n", numbers)
	// fmt.Println("numbers has length:", len(numbers))


	// create aslice with make function

	 numbers := make([]int, 3, 5)
     numbers = append(numbers, 1, 2, 3)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}