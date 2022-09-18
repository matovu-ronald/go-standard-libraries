package main

import "fmt"

func main() {
	fmt.Print("This prints without adding a new line")
	fmt.Println("This prints and adds a new line")

	sum := 20
	fmt.Println("This is sample number formatting option", sum)
	const a, b, c = 5, 5, 10
	fmt.Println("Add", a, "and", b, "to get", c)

	// Printing a slice of values
	items := []int{10, 20, 30, 40, 50}
	length, err := fmt.Println(items)
	fmt.Println(length, err)
}
