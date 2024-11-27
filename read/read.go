/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// max Length constant
const (
	maxLength = 20
)

// Struct to hold first and last name
type Person struct {
	fname string
	lname string
}

// Function to take too long names and "cut" them to maxLength chars
func maxNameLength(s string) string {
	retString := []rune(s)
	return string(retString[0:maxLength])
}

func main() {
	// Declare some variables...
	var file_name string
	var struct_names Person
	slice_of_structs := make([]Person, 0)

	// Ask user for the file name -> to work it has to be name+extention (f.e.: test.txt)
	fmt.Println("Please type in the name of the file: ")
	fmt.Scanln(&file_name)
	// Open file
	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Read file
	read_file := bufio.NewReader(file)

	for {
		line, _, err := read_file.ReadLine()
		// read if not empty (else break and end the loop)
		if len(line) > 0 {
			// split string, shorten if needed and assign to fname and lname
			names := strings.Split(string(line), " ")

			if len(names[0]) > 20 {
				names[0] = maxNameLength(names[0])
			}

			if len(names[1]) > 20 {
				names[1] = maxNameLength(names[1])
			}

			struct_names.fname, struct_names.lname = names[0], names[1]
			//fmt.Println(struct_names)
			// append struct to slice
			slice_of_structs = append(slice_of_structs, struct_names)
		}

		if err != nil {
			break
		}
	}
	//fmt.Println(slice_of_structs)
	// and finally print out the names.
	fmt.Println()
	for _, i := range slice_of_structs {
		fmt.Println(i.fname, i.lname)
	}
}
