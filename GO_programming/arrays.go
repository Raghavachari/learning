package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	// If you want to get a single item
	fmt.Println(colors[2])

	// Create an array and assign its values in a single line
	var numbers = [5]int{5, 4, 3, 2, 1}
	fmt.Println(numbers)

	//find the number of elements in an arry
	fmt.Println("Number of colors", len(colors))
	fmt.Println("Number of numbers", len(numbers))
}

// An Array is an object, just like all the other objects in Go if you pass it to a function, a copy will be made of the array rather than reference.
// But storing data in memory is just about all you can do with simple arrays, you can't easly sort them and add/remoe items into it at run time.
// For all those features, you should package your ordered data in slices rather than arrays.
