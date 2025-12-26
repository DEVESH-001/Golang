package main

import (
	"fmt"
	"maps"
)

// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
// if key-value pair is not present in map, it returns the zero value of the value type.

func main() {
	// m := make(map[string]string) // make: creates a map, key type: string, value type: string

	// //setting an element
	// m["name"] = "Golang"
	// m["area"] = "Backend"

	// fmt.Println(m["name"], m["area"]) // getting an element by giving key

	// m := make(map[string]int)
	// m["age"] = 33
	// m["price"] = 999

	// fmt.Println(m["phone"], m["price"])
	// fmt.Println("lenght of the map:->", len(m))

	// //deleting

	// delete(m, "price")
	// fmt.Println(m)
	// clear(m)
	// fmt.Println("Clearing map :->", m)

	// map without using make fnc
	// m := map[string]int{"price": 30, "phone": 3}
	// fmt.Println(m)

	//getting ele in map
	// m := map[string]int{"price": 30, "phone": 3}
	// v, ok := m["sda"] // v-> gets value, ok-> checks if key is present or not
	// fmt.Println(v)

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }
	// fmt.Println(m)

	m1 := map[string]int{"price": 30, "phone": 3}
	m2 := map[string]int{"price": 330, "phone": 33}

	fmt.Println(maps.Equal(m1,m2))
}
