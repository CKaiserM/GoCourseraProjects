/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wait_gr sync.WaitGroup
var mutex sync.Mutex

// split and convert user input into int array. Returns unsorted int slice

func ConvertInputToIntSlice(input_string string) []int {
	string_slice := strings.Split(input_string, " ")
	var int_slice = make([]int, len(string_slice))

	for idx, i := range string_slice {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		int_slice[idx] = j
	}
	return int_slice
}

// sort int slice goroutine. Added mutex to avoid race condition. Returns sorted slice
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
func SortIntSlice(unsorted_slice []int, subarray int) []int {
	mutex.Lock()
	fmt.Println("unsorted subarray", subarray, ":", unsorted_slice)
	sorted_slice := unsorted_slice
	fmt.Println("sorted subarray", subarray, ":", sorted_slice)
	mutex.Unlock()
	wait_gr.Done()
	return sorted_slice
}

// Merge 4 slices and sort ints.

func MergeAndSortSlices(p1 []int, p2 []int, p3 []int, p4 []int) []int {
	merged_slice := []int{}
	merged_slice = append(p1, p2...)
	merged_slice = append(merged_slice, p3...)
	merged_slice = append(merged_slice, p4...)
	sort.Ints(merged_slice)
	return merged_slice
}

func main() {
	var input_string string
	var slice_part int
	// Prompt the user to input a series of integers
	fmt.Println("Please provide programm with space separated integers (at least one int)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_string = scanner.Text()
	// Convert string to array
	input_int_slice := ConvertInputToIntSlice(input_string)
	////fmt.Println(input_int_array)

	// Get length and then divide
	input_slice_len := len(input_int_slice)
	////fmt.Println(input_slice_len)
	slice_part = input_slice_len / 4
	// if len > 0, partition the array into 4 approximately equal size parts
	if input_slice_len > 0 {
		slice_chunk_1 := input_int_slice[:slice_part]
		slice_chunk_2 := input_int_slice[slice_part:(2 * slice_part)]
		slice_chunk_3 := input_int_slice[(2 * slice_part):(3 * slice_part)]
		slice_chunk_4 := input_int_slice[(3 * slice_part):]

		wait_gr.Add(4)
		go SortIntSlice(slice_chunk_1, 1)
		go SortIntSlice(slice_chunk_2, 2)
		go SortIntSlice(slice_chunk_3, 3)
		go SortIntSlice(slice_chunk_4, 4)

		wait_gr.Wait()
		// Merge slices
		merged_slice := MergeAndSortSlices(slice_chunk_1, slice_chunk_2, slice_chunk_3, slice_chunk_4)
		// When sorting is complete, the main goroutine should print the entire sorted list.
		fmt.Println("Merged and sorted slice:")
		fmt.Println(merged_slice)

	} else {
		fmt.Println("insufitient number of ints")
	}

}
