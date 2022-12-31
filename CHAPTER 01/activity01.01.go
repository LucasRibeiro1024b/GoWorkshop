// Create a variable for the following:
// 	First name as a string
// 	Family name as a string
// 	Age as an int
// 	Peanut allergy as a bool
// Ensure they have an initial value.
// Print the values to the console.

package main

import (
	"fmt"
)

func main() {
	var FirstName string = "Lucas"
	var FamilyName string = "Ribeiro"
	var Age int = 23
	var PeanutAllergy bool = false

	fmt.Println(FirstName)
	fmt.Println(FamilyName)
	fmt.Println(Age)
	fmt.Println(PeanutAllergy)
}