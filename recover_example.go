package main

import (
	"fmt"
)

func crazyInc(i int) int {
	if i%42 == 0 {
		panic("The Meaning of Life found!!!!")
	}
	return i + 1
}

func counter() {
	defer func() {
		fmt.Println("Do not panic!!! its all good !")
		p := recover()
		fmt.Printf("panic: %#v\n", p)
	}()
	for i := 1; i < 50; i++ {
		fmt.Println("crazyInc:", crazyInc(i))
	}
}

func main() {
	counter()
	fmt.Println("And the liofe goes on ...")

}