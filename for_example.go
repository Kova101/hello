package main

import (
	"fmt"
	"runtime"
)

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println("Dobrogi")
	}

	j := 0
	loop := true

	for loop {
		fmt.Println("Dobrogi", j)
		j++

		if j == 4 {
			loop = false
		}
	}

	k := 0

	for {
		fmt.Println("Dobrogi", k)
		k++

		if k == 4 {
			break
		}
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Something which is not an OS.")
	}

}
