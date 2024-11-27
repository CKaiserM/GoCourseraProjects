package main

import (
	"fmt"
)

func BubbleSort(array []int) []int {

	arrLen := len(array)

	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if array[j] > array[j+1] {
				Swap(array, j, j+1)
			}
		}
	}
	return array
}

func Swap(array []int, i int, j int) {
	var temp int = array[i]
	array[i] = array[j]
	array[j] = temp
}

func main() {

	array := []int{}
	var user_input int
	fmt.Println("Please enter 10 integers")
	fmt.Println("after each int, press enter")
	for i := 0; i < 10; i++ {
		fmt.Scanln(&user_input)
		array = append(array, user_input)
	}
	fmt.Println("Unsorted Array:")
	fmt.Println(array)
	fmt.Println("Bubble Sorted Array:")
	fmt.Println(BubbleSort(array))

}
