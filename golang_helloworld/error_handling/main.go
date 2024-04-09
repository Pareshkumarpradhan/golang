package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println(("started Error Handling in the functions"))
	// ans,err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error occurred: ", err)
		
	// }
	// fmt.Println("Division of two number is :", ans)

	ans, _ := divide(10, 0)
	
	fmt.Println("Division of two number is :", ans)
}