package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["Paresh"] = 92
	studentGrades["Raj"] = 85
	studentGrades["Rahul"] = 78
	studentGrades["Ravi"] = 65

	fmt.Println("Marks of Paresh is :", studentGrades["Paresh"])
	studentGrades["Paresh"] = 95
	fmt.Println("New Marks of Paresh is :", studentGrades["Paresh"])


	delete(studentGrades, "Ravi")
	fmt.Println("Marks of Ravi is :", studentGrades["Ravi"])

	// Check if key exists
	grades, exists := studentGrades["David"]
	fmt.Println("Marks of David is :", grades, "Exists:", exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"Paresh": 92,
		"Raj": 85,
		"Rahul": 78,
		"Ravi": 65,
	}

	for index, value := range person {
		fmt.Printf("--------Key is %s and marks is %d\n", index, value)
	}
}