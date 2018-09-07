package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		os.Exit(1)
	}

	bs2 := bs[0 : len(bs)-1]

	advent1(bs2)
	advent1_2(bs2)
}

func advent1(bs []byte) {
	count := 0

	for i, v := range bs {
		ni := (i + 1) % len(bs)

		if v == bs[ni] {
			count += int(v - '0')
		}
	}

	fmt.Println(count)
}

func advent1_2(bs []byte) {
	count := 0

	if len(bs)%2 != 0 {
		os.Exit(1)
	}

	for i, v := range bs {
		ni := (i + len(bs)/2) % len(bs)

		if v == bs[ni] {
			count += int(v - '0')
		}
	}

	fmt.Println(count)
}
