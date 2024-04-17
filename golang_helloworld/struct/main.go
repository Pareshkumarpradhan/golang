package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	Email   string
	Phone   int
}

type Adderss struct {
	House int
	Area  string
	State string
}
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Adderss Adderss
}

func main() {
	var paresh Person

	paresh.FirstName = "Paresh"
	paresh.LastName = "Chouhan"
	paresh.Age = 25

	//fmt.Println(paresh.FirstName)
	//fmt.Println("Paresh persion: ", paresh)

	// 2nd method
	persion1 := Person{
		FirstName: "Paresh",
		LastName: "Kumar",
		Age: 25,
	}

	fmt.Println("Persion1: ", persion1)

	// 3rd method
	var persion2 = new(Person)
	persion2.FirstName = "simran"
	persion2.LastName = "kumar"
	persion2.Age = 25

	//fmt.Println("Persion2: ", persion2)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Paresh",
		LastName: "Chouhan",
		Age: 25,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = 1234567890

	employee1.Person_Adderss = Adderss{
		House: 123,
		Area:  "Gandhinagar",
		State: "Gujarat",
	}

	fmt.Println("Employee1: ", employee1)

}
