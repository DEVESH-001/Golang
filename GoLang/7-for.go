package main

// for -> only loop available in golang
func main() {
	//while loop
	//i := 1
	// for i <= 10 {
	// 	println(i)
	// 	i++
	// }

	//infinite loop
	// for{
	// 	println("infinite loop")
	// }

	//classic for loop

	// for i := 0; i < 20; i++ {
	// 	if i ==2{
	// 		continue
	// 	}
	// 	println(i)
	// }

	//range: used to iterate over arrays, slices, maps, strings

	for i := range 9 { //9 will be excluded because range starts from 0 to n-1
		println(i)
	}
}
