package main

import "fmt"

func main() {
	a := 42           // int
	b := 3.142        // float64
	c := 0.867 + 0.5i // complex128

	fmt.Printf("a is of type %t\n", a)
	fmt.Printf("b is of type %t\n", b)
	fmt.Printf("c is of type %t\n", c)
}
