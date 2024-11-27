package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//create name and address vars
	var input_name, input_address string
	//ask for name and address
	fmt.Println("Enter Name: ")
	fmt.Scanln(&input_name)
	fmt.Println("Enter Address: ")
	fmt.Scanln(&input_address)
	// create map from given inputs
	map_name_address := map[string]string{"name": input_name, "address": input_address}
	// and create json from map
	jn, err := json.Marshal(map_name_address)
	if err == nil {
		// Print result
		fmt.Println("JSON: ", string(jn))
	}

}
