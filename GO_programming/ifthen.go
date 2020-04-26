package main

import "fmt"

/*
Very similar to Java and C, but no parentheses around the conditions, and you have the option of adding that initial statement within the conditional block.
*/

func main() {

	var x float64 = 0
	var result string

	if x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than or equal to zero"
	}

	fmt.Println("Result: ", result)

	/*Adding initial statement within the conditional block. From below x value is applicable only in if block
	if x := 42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than or equal to zero"
	}
	*/

}
