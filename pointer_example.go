package main

import "fmt"

var pp *int
var ppp = new(int)

func main() {
	fmt.Printf("Empty pointer: %v\n", pp)

	x := 42
	p := &x
	fmt.Printf("Pointed value: %v\n", *p)

	*p = 43
	fmt.Printf("Pointed value: %v\n", *p)
	fmt.Printf("Pointer memory address: %v\n", p)

	// Modifier example
	y := 5
	fmt.Println("y:", y)

	// Technically it is not wrong, but in this example we want to use pointers.
	// If you assign the return value of the function to the variable then it works.
	wrongMultiplyByTwo(y)
	fmt.Println("Wrong multiplied y:", y)

	// Mulriply value by using pointers.
	multiplyByTwo(&y)
	fmt.Println("Good multiplied y:", y)

	// Pointer create with new()
	fmt.Println("new() pointer:", ppp, *ppp)

	np := newInt()
	fmt.Println("newInt() pointer:", np, *np)

	// Advanced
	var i *int

	if i == nil {
		i = new(int)
		*i = 3
		fmt.Println("meh")
	}

	fmt.Println(*i)
}

// Wrong method for pointer override
func wrongMultiplyByTwo(a int) int {
	return a * 2
}

// Good method
func multiplyByTwo(a *int) {
	*a *= 2
}

func newInt() *int {
	var dummy int
	return &dummy
}
