package main

import "fmt"

// powerSeries gets the power series of integer a and returns a tuple containing square of a and cube of a.
func powerSeries(a int) (square, cube int) {

	square = a*a
	cube = square*a
	return
}

func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Square: ", square,  "\nCube: ", cube)
}
