package main

import (
	"fmt"
	"sort"
)

var m map[string]string

func main() {
	m = make(map[string]string)

	m["boborján"] = "cékla"
	m["lalyos"] = "maradék"
	m["andor"] = "húsleves"
	m["peti"] = "cékla"
	m["zoli"] = "mindent"
	m["atti"] = "hambi"
	m["roli"] = "pizza"

	for k, v := range m {
		fmt.Printf("%v: %#v\n", k, v)
	}

	// Sort by key
	sortprettyMap(m)
	fmt.Printf("%#v\n", m)

	m2 := map[string]string{
		"boborján": "cékla",
		"lalyos":   "maradék",
		"andor":    "húsleves",
		"peti":     "cékla",
		"zoli":     "mindent",
		"atti":     "hambi",
		"roli":     "pizza",
	}

	fmt.Printf("%#v\n", m2)

	// Delete example
	delete(m, "boborján")

	fmt.Printf("%#v\n", m)

	favFood("andor")

	fmt.Printf("%#v\n", occurrences([]int{1, 1, 2, 2, 2, 3, 4, 4, 4, 4, 4, 5, 6, 6}))
}

func sortprettyMap(m map[string]string) {
	keys := keys(m)
	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, m[v])
	}
}

func keys(m map[string]string) []string {
	keys := make([]string, 0, len(m))

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func favFood(name string) {
	food, ok := m[name]

	if !ok {
		fmt.Printf("Who is %s?\n", name)

		return
	}

	fmt.Printf("%s's favourite food is %s\n", name, food)
}

func occurrences(in []int) map[int]int {
	ret := map[int]int{}

	for _, v := range in {
		ret[v]++
	}

	return ret
}
