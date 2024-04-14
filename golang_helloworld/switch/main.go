package main

func main() {
	var i int = 2
	switch i {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	default:
		println("Unknown")
	}

	temperature := -25
	switch {
	case temperature < 0:
		println("Temperature is below freezing")
	case temperature >= 0 && temperature < 10:
		println("It's cold")
	case temperature >= 10 && temperature < 20:
		println("It's cool")
	case temperature >= 20 && temperature < 30:
		println("It's warm")
	default:
		println("It's hot")
	}
}
