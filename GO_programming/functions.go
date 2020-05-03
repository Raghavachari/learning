package main

import "fmt"

/*
1. Go is organized with packages and packages have functions. Your own application has its own package, always named main, and it also has the main function which is called automatically by the run time as the application starts up. You can define your own custom functions and you can then organize them into their own packages
2. As long as the function definitions and the code that calls them are in the same package, the case of the function name can be either upper or lower.
3. Your functions can also receive arguments, as many arguments as you need, and each of them can have its own type. Remember there's no function overloading so each function name has to have a specific number of arguments. You can't have two versions of the function with different numbers or types of arguments.
4. You can also create functions that accept arbitrary numbers of values as long as they're all of the same type. You declare the name of the values you're passing in, then three dots and then the type.
*/

func main() {
	doSomething()
	sum := addValues(2, 3)
	fmt.Println("Sum :", sum)

	sum = addAllValues(2, 3, 4, 5)
	fmt.Println("Sum :", sum)
}

func doSomething() {
	fmt.Println("Doing something")
}

//func addValues(value1 int, value2 int) int {
func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}
