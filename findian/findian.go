/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be case-sensitive,
so it does not matter if the characters are upper-case or lower-case.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var stdin *bufio.Scanner

	var user_input string

	//Getting input from user
	fmt.Println("Please enter string:")
	stdin = bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		user_input = stdin.Text()
		break
	}
	//Getting rid of case sensitivity...
	lower_string := strings.ToLower(string(user_input[:len(user_input)]))
	//Remove whitespaces...
	ns_string := strings.ReplaceAll(lower_string, " ", "")

	// if first is "i", contains "a" and ends with "n" print "Found!", else "Not Found!"
	if strings.HasPrefix(string(ns_string), "i") && strings.Contains(string(ns_string), "a") && strings.HasSuffix(ns_string, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
