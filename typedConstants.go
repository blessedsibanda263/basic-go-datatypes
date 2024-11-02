package main

import "fmt"

const (
	typedConstant   = int16(100)
	untypedConstant = 100
)

func main() {
	i := int(1)
	fmt.Printf("unTyped:", i*untypedConstant)
	fmt.Printf("Typed:", i*typedConstant)
}
