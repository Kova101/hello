package main

import (
	"fmt"
	"time"
)

var life int = 42

func init() {
	fmt.Println("init() called")
}

func main() {
	t := time.Now()

	fmt.Println("Ahoy")
	fmt.Println(life)
	fmt.Println(t.Hour())

	// Rare - not advisable to use
	a, b := 1, "jeah"

	c, d := 1, 2

	// Tuple assignment
	c, d = d, c

	fmt.Printf("a: %v, b: %v\n", a, b)
	fmt.Printf("type - a: %T, b: %T\n", a, b)
	fmt.Printf("c: %v, d: %v\n", c, d)
}
