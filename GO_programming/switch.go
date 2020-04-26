package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Go's switch statement expands on C's and Java's versions. Like the most recent versions of Java you can evaluate any simple type, not just constants or numerics and there's more flexibility in the syntax compared to both of those ancestral languages. Also in Go, Code Flow automatically jumps past additional cases after finding a match. So you don't need a bunch of redundant break statements

You don't need to add the break keyword to stop flow from falling through from a successful case to the rest of the switch statement. By default, you jump to the end of the switch statement

Go's fall-through keyword. If you add it after any code within a case if that case it true you'll execute its code, and you'll also execute the next case

The difference is between Go and C and Java's version is that the value you're evaluating doesn't get any parentheses. You can add a statement before the evaluation and any variables that you declare within that statement will be local to the switch statement and you don't need to add the break keywords within each case.
*/

func main() {

	//seeds the randomized value using a number of milliseconds from the current date and time
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
	case 1:
		result = "It's Sunday"
	case 7:
		result = "It's Saturday"
	default:
		result = "Its Weekday"
	}
	fmt.Println("Day", dow, ",", result)

	x := -42
	switch {
	case x < 0:
		result = "Less than zero"
		// fallthrough
	case x == 0:
		result = "Equal to zero"
	default:
		result = "Greater than zero"
	}
	fmt.Println(result)

}
