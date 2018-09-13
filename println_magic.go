package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("value:%v type:%T address:%p\n", s, s, s)
	// reusing argument
	fmt.Printf("value:%[1]v type:%[1]T address:%[1]p\n", s)
	fmt.Printf("short pi:%.2f\n", 3.1415926)
	fmt.Printf("right paddin gstring: %10s\n", "short")
	fmt.Printf("left paddin gstring: %-10s\n", "short")
	// truncating string, is it documented ?!
	fmt.Printf("sort string: %10.10s\n", "whatifastrinisreallyreallylong")
	fmt.Printf("trunc pi by var:%.[2]*[1]f\n", 3.1415926, 4)
	fmt.Printf("truncating str by var: %[1]*.[1]*[2]s\n", 4, "whatifastrinisreallyreallylong")
}