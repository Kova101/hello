package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Group []Person

type Sorting func(i, j int) bool

func (g Group) SortBy(t string) Sorting {
	return func(i, j int) bool {
		switch t {
		case "name":
			return g[i].Name < g[j].Name
		case "age":
			return g[i].Age < g[j].Age
		}

		return false
	}
}

func groupByString(g Group, f func(p Person)string) Sorting {
	return func(i, j int) bool {
		return f(g[i]) < f(g[j])
	}
}

func main() {
	group1 := Group{
		Person{"adel", 33},
		Person{"zoe", 18},
		Person{"bob", 22},
		Person{"rob", 45},
	}
	group2 := Group{
		Person{"jeff", 39},
		Person{"doroty", 58},
		Person{"maria", 9},
		Person{"george", 101},
	}

	sort.SliceStable(group1, group1.SortBy("name"))
	fmt.Println("group1 byName", group1)

	sort.SliceStable(group1, group1.SortBy("age"))
	fmt.Println("group1 byAge", group1)

	sort.SliceStable(group2, group2.SortBy("name"))
	fmt.Println("group2 byName", group2)

	sort.SliceStable(group2, group2.SortBy("age"))
	fmt.Println("group2 byAge", group2)

	sort.SliceStable(group2, groupByString(group2, func(p Person)string {return p.Name}))
	fmt.Println("group2 byAge", group2)

}
