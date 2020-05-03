package main

import "fmt"

func main() {
	n1, l1 := FullName("Raghava", "Chari")
	fmt.Printf("FullName: %v, number of characters: %v\n\n", n1, l1)

	n2, l2 := FullNameNakedReturn("Raghava", "Chari")
	fmt.Printf("FullName: %v, number of characters: %v\n\n", n2, l2)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
