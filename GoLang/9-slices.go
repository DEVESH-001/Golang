package main

import (
	"fmt"
	"slices"
)

//slices -> are dynamic arrays, can grow and shrink in size
// mostly used + more powerful than arrays

func main() {

	//var nums []int // declaring a slice [uninitialized]
	//vs arrays: var nums [5]int

	//fmt.Println(nums)
	//fmt.Println(len(nums))

	// give initial values
	// var nums = make([]int, 0, 5) // len: 2, cap: 5, make function to create slice
	// fmt.Println(cap(nums))       // capacity: max nums of elements slice can hold
	// //fmt.Println(nums)

	// nums = append(nums, 1) // append function to add elements to slice
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(cap(nums)) // capacity increases as we append more elements
	// nums = append(nums, 10)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// nums := []int{} // shorthand to declare and initialize an empty slice
	// nums[0] = 2
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// Copy fnc

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	// Slice:  gives a more powerful interface to sequences than arrays.
	var nums = []int{1, 2, 3, 4, 5}

	fmt.Println(nums[:3]) // slicing syntax: [startIndex : endIndex(exclusive)]

	var nums1 = []int{1, 2, 3}
	var nums2 = []int{1, 2, 3}

	fmt.Println("Comparing 2 :->", slices.Equal(nums1, nums2))

	// 2D Slices

	var nums3 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Printing 2Dslices:->", nums3)
}
