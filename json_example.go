package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City   string
	Street string `json:"address"`
}

type Person struct {
	Address
	Name string
	Age  int
}

var listJson = `
[
  {
    "address": {"city": "budapest","address": "Bajcsy"},
    "name": "joe",
    "age": 33
  },
  {
    "city": "new york",
    "address": "Bajcsy",
    "name": "bill",
    "age": 44
  }
]
`

func main() {
	var p Person

	p = Person{
		Name: "joe",
		Age:  33,
		Address: Address{
			City:   "budapest",
			Street: "Bajcsy",
		},
	}

	fmt.Printf("%#v\n", p)

	var g []Person

	json.Unmarshal([]byte(listJson), &g)

	fmt.Println("===> unmarshal")

	for _,p := range g {
		fmt.Printf("%#v\n", p)
	}
}
