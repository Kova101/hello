package main

import "fmt"

var a = [5]int{1, 2, 3, 4, 5}
var b = [...]int{1, 2, 3}
var name = "Géza"
var slice []int
var sliceA []int

func main() {
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	for k, v := range b {
		fmt.Println(k, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	// String byte examples
	fmt.Println("Name byte count:", len(name))

	for _, v := range name {
		fmt.Printf("%c\n", v)
	}

	var bs []byte
	bs = []byte(name)
	rs := []rune(name)

	fmt.Println(name, "====", bs)
	fmt.Println(name, "====", rs)

	// Multi-dimension
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println("======== Array examples (slices) ========")

	// Slice example
	slice = a[2:3]
	fmt.Printf("slice: %#v %T\n", slice, slice)
	fmt.Printf("Nil?: %b Len: %d, Capacity: %d\n", slice == nil, len(slice), cap(slice))

	slice2 := slice[0:cap(slice)]
	fmt.Println("Slice2:", slice2)

	// Slice example array's first N square value
	sliceA = a[0:1]
	fmt.Println(a)
	sInfo(sliceA)
	for i := 0; i < 10; i++ {
		sliceA = append(sliceA, 1)
		sInfo(sliceA)
	}
	fmt.Println(a)
}

func sInfo(s []int) {
	fmt.Println(s, len(s), cap(s), s == nil)
}
