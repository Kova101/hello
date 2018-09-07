package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1 = make([]int, 0, 5)
	var s2 = []int{6, 4, 7, 8, 3}
	var names = []string{"Atti", "Roli", "Peti", "Andor", "Lalyos"}

	fmt.Printf("%#v\n", names)

	sort.Strings(names)

	fmt.Printf("%#v\n", names)
	sinfo(s1)
	sinfo(s2)
}

func ainfo(a [5]int) {
	fmt.Printf("%#v\n", a)
}

func sinfo(s []int) {
	fmt.Printf("%#v len: %v, cap: %v\n", s, len(s), cap(s))
}
