package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var b int
var v string
var n bool
var s bool

func init() {
	flag.IntVar(&b, "beat", 3, "How many times to beat")
	flag.StringVar(&v, "victim", "Döbrögi", "Who to beat")
	flag.BoolVarP(&n, "number", "n", false, "Count beats")
	flag.BoolVarP(&s, "stick", "s", false, "Beat with stick")

	flag.Parse()
}

func main() {
	fmt.Println(b)
	fmt.Println(v)

	fmt.Printf("Beat %v, %v times!\n", v, b)

	for i := 0; i < b; i++ {
		if n {
			fmt.Print(i+1, " ", v)
		} else {
			fmt.Print(v)
		}

		if s {
			fmt.Print(" [STICK]")
		}
		fmt.Println()
	}
}
