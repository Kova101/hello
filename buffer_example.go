package main

import (
	"bytes"
	"fmt"
)

func main() {
	s1 := "hello"

	prettySlice([]byte(s1))
	prettySlice([]byte("Árvíztűrő tükörfúrógép"))
}

func prettySlice(bs []byte) {
	var b1, b2, b3 bytes.Buffer

	for k, v := range bs {
		fmt.Fprintf(&b1, "%3d.|", k)
		fmt.Fprintf(&b2, "%#x|", v)
		fmt.Fprintf(&b3, "%4c|", v)
	}

	fmt.Println(b1.String())
	fmt.Println(b2.String())
	fmt.Println(b3.String())
}
