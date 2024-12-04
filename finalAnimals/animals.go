/*
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomotion method	Spoken sound
cow		grass 		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake
all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
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

// Bird actions
func (e Bird) Eat() {
	fmt.Println(e.animal_name, "eats warms")
}

func (m Bird) Move() {
	fmt.Println(m.animal_name, "flys")
}

func (s Bird) Speak() {
	fmt.Println(s.animal_name, "peeps")
}

// Cow actions
func (e Cow) Eat() {
	fmt.Println(e.animal_name, "eats grass")
}

func (m Cow) Move() {
	fmt.Println(m.animal_name, "walks")
}

func (s Cow) Speak() {
	fmt.Println(s.animal_name, "moos")
}

// Snake actions
func (e Snake) Eat() {
	fmt.Println(e.animal_name, "eats mice")
}

func (m Snake) Move() {
	fmt.Println(m.animal_name, "slithers")
}

func (s Snake) Speak() {
	fmt.Println(s.animal_name, "hisses")
}

// function creates new cow, bird or snake with given name (append animal)
func NewAnimal(animals_slice *[]Animal, animal_type string, animal_name string) {
	if animal_type == "bird" {
		*animals_slice = append(*animals_slice, Bird{animal_name})
		fmt.Println("Created it!")
	} else if animal_type == "cow" {
		*animals_slice = append(*animals_slice, Cow{animal_name})
		fmt.Println("Created it!")
	} else if animal_type == "snake" {
		*animals_slice = append(*animals_slice, Snake{animal_name})
		fmt.Println("Created it!")
	} else {
		fmt.Println(animal_type, "is not a cow, bird or snake.")
	}
}

// function returns animal query (animal name+action)
func GetAnimalQuery(animal_slice []Animal, get_animal_name string, get_animal_action string) {
	for _, a := range animal_slice {
		cow_value, ok := a.(Cow)
		if ok {
			if get_animal_name == cow_value.animal_name {
				GetAnimalAction(cow_value, get_animal_action)
			}
		}
		snake_value, ok := a.(Snake)
		if ok {
			if get_animal_name == snake_value.animal_name {
				GetAnimalAction(snake_value, get_animal_action)
			}
		}
		bird_value, ok := a.(Bird)
		if ok {
			if get_animal_name == bird_value.animal_name {
				GetAnimalAction(bird_value, get_animal_action)
			}
		}
	}
}

// function returns animal action
func GetAnimalAction(animal Animal, animal_action string) {
	switch animal_action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
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
		//get user input
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_command := scanner.Text()
		//split user commands
		input_slice = strings.Split(input_command, " ")
		input_cmd := strings.ToLower(input_slice[0])
		//fmt.Println(input_slice)
		if input_cmd == "newanimal" {
			input_animal_name := input_slice[1]
			input_animal_type := strings.ToLower(input_slice[2])
			//add new animal
			NewAnimal(&animals_slice, input_animal_type, input_animal_name)
		} else if input_cmd == "query" {
			input_animal_name := input_slice[1]
			input_animal_action := strings.ToLower(input_slice[2])
			//get user query
			GetAnimalQuery(animals_slice, input_animal_name, input_animal_action)
		} else {
			//invalid input
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
