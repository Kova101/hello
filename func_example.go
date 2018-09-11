package main

import "fmt"

var Roland []string

type Shoe struct {
	Id     int
	Model  string
	Size   []int
	Brand  string
	Price  float64
	Wproof bool
}

type SrcCond func(Shoe) bool

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Sum of %v\n", sum(s1...))
	fmt.Printf("Sum of %v\n", sum(s2...))
	fmt.Printf("Cap: %v\n", cap(Roland))

	r, _ := findSomething(10)
	fmt.Println(r)


	_, err := findSomething(11)

	if err != nil {
		fmt.Println(err)
	}
}

func sum(ns ...int) int {
	sum := 0

	for _, v := range ns {
		sum += v
	}

	return sum
}

func findSomething(i int) (int, error) {
	if i == 10 {
		return i, nil
	}

	return 0, fmt.Errorf("not 10 sorry")
}

func findShoe(b string) SrcCond {
	return func(s Shoe) bool {
		return s.Brand == b
	}
}
