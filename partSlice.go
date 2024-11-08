package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	// First 5 elements
	fmt.Println(aSlice[0:5])
	// First 5 elements
	fmt.Println(aSlice[:5])

	l := len(aSlice)
	// Last 2 elements
	fmt.Println(aSlice[l-2 : l])
	// Last 2 elements
	fmt.Println(aSlice[l-2:])

	// First 5 elements
	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))

	// Elements at indexes 2,3,4
	t = aSlice[2:5:10] // capacity will be 10-2=8
	fmt.Println(len(t), cap(t))

	// Elements at indexes 0,1,2,3,4
	t = aSlice[:5:6] // New capacity will be 6-0
	fmt.Println(len(t), cap(t))
}
