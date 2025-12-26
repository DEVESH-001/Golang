// In Go, an array is a numbered sequence of elements of a specific length.
// Arrays have a fixed size, which means that once an array is created, its size cannot be changed.
// Arrays can hold elements of any type, including built-in types like int, float64, and string

package main

func main() {
	//var nums [4]int // Declare an array of integers with a length of 4

	//arr length
	//println(len(nums))

	//adding values to array
	// nums[0] = 199
	// fmt.Println(nums[0])
	// fmt.Println(nums)

	// var vals [4]string // Declare an array of booleans with a length of 4
	// vals[0] = "true"
	// vals[2] = "sher"

	// fmt.Println(vals)

	//adding values while declaring
	// nums := [3]int{1, 22, 3}
	// fmt.Println(nums)

	// 2D arrays
	// nums := [2][2]int{{11, 2}, {3, 44}}
	// fmt.Println(nums)
}

//when to use arrays?
// Arrays are useful when you know the exact number of elements you need to store and that number will not change.
// They provide fast access to elements based on their index, making them suitable for scenarios where performance is critical.
// Memory optimization is another reason to use arrays, as they have a fixed size and can be more memory-efficient than other data structures like slices or lists.
