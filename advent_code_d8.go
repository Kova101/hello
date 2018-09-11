package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands = map[string]func(in int, f string){}
var r = map[string]int{}

// Collect the highest value ever on a register
var highVal = 0

func init() {
	commands["INC"] = func(in int, f string) {
		r[f] += in

		if r[f] > highVal {
			highVal = r[f]
		}
	}
	commands["DEC"] = func(in int, f string) {
		r[f] -= in

		if r[f] > highVal {
			highVal = r[f]
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan file line by line
	for scanner.Scan() {
		t := scanner.Text()

		// Split line into command and statement
		l := strings.Split(t, " if ")

		runCommands(l)
	}

	// Get the max value from the register map
	var max = 0

	for _, v := range r {
		if v > max {
			max = v
		}
	}

	fmt.Printf("%v\n", max)
	fmt.Printf("%v\n", highVal)
}

func runCommands(l []string) {
	// Split the command and the statement into words and check the slice length
	s := strings.Fields(l[0])

	if len(s) != 3  {
		return
	}

	statement := strings.Fields(l[1])

	if len(statement) != 3  {
		return
	}

	// Assign command values
	v, c, n := s[0], s[1], s[2]

	// Get command type from command map
	com, ok := commands[strings.ToUpper(c)]

	if !ok {
		return
	}

	// Convert command number string to integer
	num, err := strconv.Atoi(n)

	if err != nil {
		return
	}

	// Check if we should do the command
	if !doCommand(statement) {
		return
	}

	com(num, v)
}

func doCommand(s []string) bool {
	// Get the register value from the register map
	i := r[s[0]]

	// Convert the command numeric string to integer
	num, err := strconv.Atoi(s[2])

	if err != nil {
		return false
	}

	// Use the correct operation on the two integers
	switch s[1] {
	case ">":
		return i > num
	case "<":
		return i < num
	case "==":
		return i == num
	case "!=":
		return i != num
	case "<=":
		return i <= num
	case ">=":
		return i >= num
	}


	return false
}
