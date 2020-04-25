package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		fmt.Println("Value of p is", *p)
	} else {
		fmt.Println("P is nil")
	}

	var v int = 42
	// point the pointer at the value v. & means connect the pointer to this variable
	p = &v
	if p != nil {
		fmt.Println("Value of p is", *p)
	} else {
		fmt.Println("P is nil")
	}

	// reference the pointer value using * operator
	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("Value1:", *pointer1)

	// modify the pointer value
	*pointer1 = *pointer1 / 31
	fmt.Println("Value1:", *pointer1)
	fmt.Println("Value1:", value1)

}
