package main

import "fmt"

func main() {
	// Byte slice
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)
	fmt.Println(len(b), cap(b))

	// Print byte slice contents as text
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))
}
