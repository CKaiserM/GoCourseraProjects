/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal:

1) the food that it eats,
2) its method of locomotion, and
3) the sound it makes when it speaks.

The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion method	Spoken sound
cow		grass 		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user,
prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.

Submit your Go program source code.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
*/
type Animal interface {
	Eat()
	Move()
	Speak()
}

/*
Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
*/

type Bird struct{ animal_name string }
type Cow struct{ animal_name string }
type Snake struct{ animal_name string }

func (e Bird) Eat() {
	fmt.Println("warms")
}

func (m Bird) Move() {
	fmt.Println("fly")
}

func (s Bird) Speak() {
	fmt.Println("peep")
}

func (e Cow) Eat() {
	fmt.Println("grass")
}

func (m Cow) Move() {
	fmt.Println("walk")
}

func (s Cow) Speak() {
	fmt.Println("moo")
}

func (e Snake) Eat() {
	fmt.Println("mice")
}

func (m Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// input variables
	var decision string
	input_slice := []string{}
	animals_slice := []Animal{}

	fmt.Println("To create new animal, enter <newanimal> command following with <animal name> and the string, either “cow”, “bird”, or “snake”.")
	fmt.Println("or to run query enter <query> <animal name> and the string, either “eat”, “move”, or “speak”.")
	// start interface loop

interfaceLoop:
	for {

		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_command := scanner.Text()

		input_slice = strings.Split(input_command, " ")

		fmt.Println(input_slice)
		if strings.ToLower(input_slice[0]) == "newanimal" {
			switch strings.ToLower(input_slice[2]) {
			case "bird":
				animals_slice = append(animals_slice, Bird{animal_name: input_slice[1]})
				fmt.Println("Created it!")
			case "cow":
				animals_slice = append(animals_slice, Cow{animal_name: input_slice[1]})
				fmt.Println("Created it!")
			case "snake":
				animals_slice = append(animals_slice, Snake{animal_name: input_slice[1]})
				fmt.Println("Created it!")
			}
		} else if strings.ToLower(input_slice[0]) == "query" {
			switch strings.ToLower(input_slice[2]) {
			case "eat":

			case "move":

			case "speak":

			}
		} else {
			fmt.Println("invalid command")
			fmt.Println("to exit type 'x', to reset hit enter")
			fmt.Scan(&decision)
			if strings.ToLower(decision) == "x" {
				break interfaceLoop
			}
			continue
		}

	}

}
