package main

import (
	"fmt"
	"sort"
)

/*
1. A MAP in Go is a unordered collection of key value pair.
2. Its essentially a hashtable and stores the collection of data and then orbitarly find the items in collection by using keys.
3. A MAP's key can be of any type that can be comparable. i.e the keys can be compared to each other for the purpose of sorting.
   But its pretty common to use strings for keys and any other type for the associated values.
4. You can declare empyt map with builtin make function.
*/

func main() {
	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregan"
	states["CA"] = "California"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	// delete an item
	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "NewYork"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
	// To list the states in alphabetical order. Extract the keys from the map as slice
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	fmt.Println("\nSorted")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
