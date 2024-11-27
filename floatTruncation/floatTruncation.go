package main

import (
	"fmt"
)

func main() {

	// Declaring variable to store input
	var float_value float64

	//Getting input from user
	fmt.Println("Please enter float number:")
	fmt.Scan(&float_value)

	// cast input to int
	// Note: !=float returns 0
	fmt.Printf("%d", int(float_value))

	/*var i interface{} = float_value

	if i != 0 {
		// Printing the given texts
		fmt.Println("Truncated float:")
		fmt.Printf("%d", int(float_value))
	} else {
		fmt.Println("Error, this is not a float")
	}
	*/

}
