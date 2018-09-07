package main

import "fmt"

func main() {
	var a [5]int
	var s []int

	ainfo(a)
	a = [...]int{1, 2, 3, 4, 5}
	ainfo(a)

	sinfo(s)
	s = a[2:]
	sinfo(s)

	ainfo(a)
	arev(&a)
	ainfo(a)

	sinfo(s)
	srev(s)
	sinfo(s)
}

func ainfo(a [5]int) {
	fmt.Printf("%#v\n", a)
}

func sinfo(s []int) {
	fmt.Printf("%#v len: %v, cap: %v\n", s, len(s), cap(s))
}

func arev(a *[5]int) {
	for i := 0; i < 2; i++ {
		a[i], a[4-i] = a[4-i], a[i]
	}
}

func srev(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
