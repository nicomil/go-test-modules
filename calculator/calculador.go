package main

import (
	// "calculation" is a module loaded in system /usr/local/go/src/calculation available globally.
	// to create a global module go to /usr/local/go/src/ (mac), create a new folder with the module name, enter in the folder, create a file .go with all functions you want, compile with sudo go install
	"fmt"
	"mathOps"
	"utilities"
)

func main() {
	fact := mathOps.Fact(4)

	fmt.Printf("Fact is %d\n", fact)
	n := 10
	sumn := mathOps.SumN(n)
	fmt.Printf("Sum of the first %d numbers is %d\n", n, sumn)

	arr := []string{"Nico", "Gracia", "Nahu"}
	utilities.PrintArr(arr)

	fmt.Println(utilities.StrConcat(arr...))
}
