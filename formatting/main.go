package main

import "fmt"

func main() {
	x := 10
	y := 12.25

	// Basic formatting
	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)

	// Print Booleans using "True" or "False"
	fmt.Printf("%t\n", x > 10)

	// Print floating point and exponential numbers
	fmt.Printf("%f\n", y)
	fmt.Printf("%e\n", y)

	// Using explicit argument indexes
	fmt.Printf("%[2]d %[1]d\n", 20, 23)

	// Argument indexes can be used to print values repeatedly
	fmt.Printf("%d %#[1]o %#[1]x\n", 15)

	// Print a value in default format
	c := circle{
		radius: 7,
		border: 2,
	}

	fmt.Printf("%+v\n", c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", c)
	// Sprintf works as Printf but returns a string
	s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
	fmt.Println(s)

}

type circle struct {
	radius int
	border int
}
