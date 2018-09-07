package main

import (
	"fmt"
	"go/build"
	"os"
	"strings"
)

var path string

func main() {
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	needle := build.Default.GOPATH + string(os.PathSeparator) + "bin"
	found := false

	for _, v := range paths {
		if needle == v {
			found = true
		}
	}

	if !found {
		ss := []string{needle}
		paths = append(ss, paths...)

		fmt.Printf("export PATH=%v\n", strings.Join(paths, string(os.PathListSeparator)))
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Mehh")
}
