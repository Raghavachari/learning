package main

import (
	"fmt"
)

/*
1. The struct type in Go is a data structure. Structs are similar in purpose and function to C's struct and Java's classes.
2. They encapsulate both data and methods, but unlike Java, Go doesn't have an inheritance model. You don't have concepts like super or sub-structs.
3. Each structure is independent, with its own fields for data management and optionally its own methods.
4. You define a struct as its own custom type. The struct's name typically has an initial uppercase character. Start with type, then the struct name
5. I'll name mine Dog, and then the keyword struct. Within the braces, you can then define the fields to store data that's associated with the structure. In order to make sure that the fields are publicly accessible, start with an initial uppercase character, and then keep the names to single words that are easily understandable.
*/

type Dog struct {
	Breed  string
	Weight int
}

func main() {
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	// You can dump the complete contents of a struct including its field names and its values using Printf fnction by using verb %+v
	fmt.Printf("%+v\n", poodle)
	// Each value of a struct is known as fields like java's classes. You can access the structs individual fields using dot(.) operator
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}
