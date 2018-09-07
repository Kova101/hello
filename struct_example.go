package main

import "fmt"

// Type definition non-literal
type Person struct {
	Name string
	Age  int
}

type Group map[string]Person

var g = Group{}

func main() {
	add("Jenő", 20)
	add("Béla", 33)
	add("Cékla", 1)
	add("Répa", 88)
	list()
	remove("Béla")
	list()
	search("Répa")
	search("Non existent person")
}

func add(name string, age int) {
	fmt.Printf("Adding: %v\n", name)

	g[name] = Person{
		Name: name,
		Age:  age,
	}
}

func remove(name string) {
	fmt.Printf("Removing: %v\n", name)

	delete(g, name)
}

func list() {
	fmt.Printf("%3s. %10s (%s)\n", "ID", "NAME", "AGE")

	for i, p := range g {
		fmt.Printf("%3s. %10s (%d)\n", i, p.Name, p.Age)
	}
}

func search(name string) {
	fmt.Printf("Searching: %v\n", name)

	v, ok := g[name]

	if !ok {
		fmt.Printf("%v not found\n", name)

		return
	}

	fmt.Println(v)
}
