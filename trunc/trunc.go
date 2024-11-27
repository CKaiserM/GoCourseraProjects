package main

import (
	"fmt"
)

func main() {

	// Declaring variable to store user input
	var float_value float64

	//Getting input from user
	fmt.Println("Please enter a float number:")
	fmt.Scan(&float_value)

	// cast input to int
	// Note: !=float returns 0
	fmt.Printf("%d", int(float_value))

}

/*

package main

import (
	"fmt"
	"strconv"
)

func main() {
	sl := []string{"first", "second"}
	for _, val := range sl {
		process(val)
	}
}

func process(n string) {
	var s string
	var target float64
	fmt.Printf("Enter %s number: ", n)
	for {
		_, err := fmt.Scan(&s)
		target, err = strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("ERROR! Enter a valid %s number: ", n)
		} else {
			result := int(target)
			fmt.Printf("Result: %d\n", result)
			break
		}
	}
}

*/
