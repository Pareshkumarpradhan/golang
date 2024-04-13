package main

func main(){
	x := 10

	if x > 5 {
		println("x is greater than 5")
	} else {
		println("x is less than or equal to 5")
	}

	z := 5
	if z > 5 {
		println("z is greater than 5")
	} else if z < 5 {
		println("z is less than 5")
	} else {
		println("z is equal to 5")
	}
}