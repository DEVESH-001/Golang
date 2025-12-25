// Golang

package main // package declaration

import (
	"fmt"
	"math/rand"
) //importing a package

func main() {
	fmt.Print("Hello, World!") //using the Print function from the fmt package
	println("Hello")
	fmt.Println(rand.Int31n(10))
}
