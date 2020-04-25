package main

import (
	"fmt"
	"sort"
)

/*
A SLICE in Go is an abstration layer. It sits on top of an array. When you declare a Slice, then run time creates the
underline array for you, allocates requred memory and returns the requested Slice.
- Like Arrays all items in the slice must be of same type
- But unlike arrays it can be resizable
- It can be sorted quite easly
*/

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// Add an item to slice
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// Remove an item
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[1:])
	fmt.Println(colors)
	// To remove the last item
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 12
	numbers[2] = 13
	numbers[3] = 17
	numbers[4] = 16
	fmt.Println(numbers)

	numbers = append(numbers, 13)
	fmt.Println(numbers)

	fmt.Println(cap(numbers))
	//sort the slice
	sort.Ints(numbers)
	fmt.Println(numbers)

}
