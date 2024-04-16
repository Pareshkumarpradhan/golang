package main

import "fmt"

func main() {
	for i := 0; i < 1; i++ {
		fmt.Println("Number is :", i)
	}

	counter := 0
	for{
		fmt.Println("Counter is :", counter)
		counter++
		if counter == 1 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
         fmt.Printf("Index is %d and Value is %d\n", index, value)
	}

	data := "Hello World"
	for index, char := range data {
		fmt.Printf("Index is %d and Char is %c\n", index, char)
	}
}
