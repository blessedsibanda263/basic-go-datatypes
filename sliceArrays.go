package main

import "fmt"

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	var S0 = a[0:1] // connect S0 with 1st element of array
	fmt.Println(S0)

	var S12 = a[1:3]
	fmt.Println(S12)

	S12[0] = "S12_0"
	S12[1] = "S12_1"

	// The contents of slice `a` have changed!
	fmt.Println("a:", a)

	// changes to slice -> changes to array (because the slice and array are connected)
	change(S12)
	fmt.Println("a:", a)

	// capacity of S0
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))

	// Add 4 elements to S0
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")

	a[0] = "-N1"

	fmt.Println(S0)

	// Changing the capacity of S0
	// Not the same underlying array anymore!
	S0 = append(S0, "N4")
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))

	// This change does not go to S0
	a[0] = "-N1-"
	// This change goes to S12
	a[1] = "-N2-"
	fmt.Println("S0:", S0)
	fmt.Println("a:", a)
	fmt.Println("S12:", S12)
}
