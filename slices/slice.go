//MCK 11-25-2024

package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func capLen3() {
	/*
		Literal interpretation:
		"Before entering the loop, the program should create an empty integer slice of size (LENGTH) 3."

		Creating an empty slice [0 0 0], for the first three loops replace "0" with an int and sort the output.

	*/

	// create an "empty" slice of length 3
	sorted_slice := make([]int, 3)
	//input value
	var input_val string
	//initiate counter (used for the first 3 ints replaced by the slices.Replace*() function)
	first_three := 0

	fmt.Println("Initial slice lenght is:", len(sorted_slice), " and the capacity is: ", cap(sorted_slice))
	// start the loop
	for {
		fmt.Println("Type int to append into the slice or x to exit")
		// grab user input
		fmt.Scanln(&input_val)
		// if "x" or "X" exit the loop
		if strings.ToLower(input_val) == "x" {
			break
		}
		// if not an int, skip and ask again.
		if s, err := strconv.Atoi(input_val); err != nil {
			println(s, " is not a valid value, please enter int or x")
			continue
		}

		// lets sort the firts three: replace zeroes with the input and sort all ints.
		if first_three < 3 {
			if s, err := strconv.Atoi(input_val); err == nil {
				// if input is < 0, or slice already has negative ints in it, put in place of existing 0.
				if s < 0 || sorted_slice[0] < 0 {
					ind := slices.Index(sorted_slice, 0)
					sorted_slice = slices.Replace(sorted_slice, ind, ind+1, s)
				} else {
					// else replace first index (0) with input and sort
					sorted_slice = slices.Replace(sorted_slice, 0, 1, s)
				}

				slices.Sort(sorted_slice)
				fmt.Println(sorted_slice)
			}
			//else append to existing slice and sort
		} else {
			if s, err := strconv.Atoi(input_val); err == nil {
				sorted_slice = sort.IntSlice(append(sorted_slice, s))
				slices.Sort(sorted_slice)
				fmt.Println(sorted_slice)
			}
		}
		first_three++
	}
}

func cap3Len0() {
	/*
		Logical interpretation:
		"Before entering the loop, the program should create an empty integer slice of size (LENGTH) 3."

		Actual creating an empty slice [], but with the capacity of 3 (memory reserved for 3 ints).

		"The length of a slice is the number of elements it contains.
		The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice."

	*/

	// create empty slice of cap 3
	sorted_slice := make([]int, 0, 3)
	//input value
	var input_val string

	fmt.Println("Initial slice lenght is:", len(sorted_slice), " and the capacity is: ", cap(sorted_slice))
	//while not x, append to the slice
	for {
		fmt.Println("Type int to append into the slice or x to exit")
		fmt.Scanln(&input_val)

		if strings.ToLower(input_val) == "x" {
			break
		}
		// if not an int, skip and ask again.
		if s, err := strconv.Atoi(input_val); err != nil {
			println(s, " is not a valid value, please enter int or x")
			continue
		}
		if s, err := strconv.Atoi(input_val); err == nil {
			sorted_slice = sort.IntSlice(append(sorted_slice, s))
			slices.Sort(sorted_slice)
			fmt.Println(sorted_slice)
		}
	}
}
func main() {
	// choose your method carefully :)
	var init_input string
	fmt.Println("Please chose the interpretation of the instructions:")
	fmt.Println("For the interpretation: Lenght of 3 please type in 1; for Capacity of 3 type in 2")
	fmt.Println("OR type x to exit")
	fmt.Scanln(&init_input)

	//choose loop
	switch init_input {
	case "1":
		fmt.Println("Entering first interpretation")
		capLen3()
	case "2":
		cap3Len0()
	case "x", "X":
		break
	default:
		fmt.Println("Incorrect input")
		main()
	}
}
