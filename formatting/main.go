package main

import "fmt"

func main() {
	f := 125.3378

	// Control the precision
	fmt.Printf("%.2f\n", f)

	// Print with width e.g 10 and default precision
	fmt.Printf("%10f\n", f)

	// Print with width e.g 10, padding and specific precision e.g .2
	fmt.Printf("%10.3f\n", f)

	// Always use a + sign
	fmt.Printf("%+10.3f\n", f)

	// Pad with 0s instead of spaces
	fmt.Printf("%010.2f\n", f)

}
