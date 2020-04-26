package main

import "fmt"

/*
You can loop through your code in Go using the for statement. Unlike in C and Java, there is no while statement but you can do the same sort of iterating of your code with extended versions of for. The most common for statement is just like in other C-style languages with a three-part declaration but you drop the parentheses around those three parts.

*/
func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	//Now that traditional for loop works fine with the slice but there's a more concise and more readable version using the ////range keyword

	for i := range colors {
		fmt.Println(colors[i])
	}

	// Go doesn't have a traditional while loop but you can accomplish the same thing by using the for statement and eliminating the pre and post sections, and just evaluate a condition
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
	}

	//the break and continue statements work just like they do in other C-style languages. Break means jump to the end of the current code block and works with both for and switch statements, and continue means start back at the beginning of this code block and reevaluate any condition
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 500 {
			break
		}
	}
	//And there's also support for the good old goto statement in labels. You can add labels to your code and then jump to the label with a goto statement. So I'm going to down here, and I'll create a label named endofprogram, and I'll mark it as a label with a colon character. Then after the colon I'll say fmt.Println and then "end of program"
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}
endofprogram:
	fmt.Println("end of program")

}
