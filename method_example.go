package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/lalyos/hello/shoeshop"
)

type Person struct {
	Name string
	Age  int
}

const emailFormat = `Dear %v!

You are %d years old.

Best regards,
IOT
`

func main() {
	p := Person{
		Name: "Béla",
		Age:  33,
	}

	p.Email()

	fmt.Println("Before ageing", p)
	p.Older()
	fmt.Println("After ageing", p)

	fmt.Println("========== SHOE SHOP ==========")

	s := shoeshop.Shop{
		Name: "Decathlon",
		Address: shoeshop.Address{
			City:   "Budapest",
			Street: "Utca 1",
		},
	}

	s.Add("Comfort", "Nike", 100.5, "red", "yellow")
	s.Add("Premium", "Tisza", 130.9, "opal", "green")
	s.Add("Comfort Plus", "Nike", 120.5, "fire-red", "sun-yellow")
	s.Add("Badimass", "Adidas", 129.5, "shame-red", "jelly-yellow")
	s.Add("Badimess", "Adidas", 109.5, "sunset-red", "cheese-yellow")

	bb := s.Brands()
	fmt.Printf("Brands: %#v\n", bb)

	s.PriceList()

	b := s.FindByBrand("Nike")
	shoeshop.PrintList(b)

	s.DeleteById("ComfortNike")

	s.PriceList()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(s)

	if err != nil {
		fmt.Println("Encode error")
	}

	var s2 shoeshop.Shop
	err = dec.Decode(&s2)

	if err != nil {
		fmt.Println("Decode error")
	}

	fmt.Println("========== SHOE SHOP 2 gob ==========")
	s2.PriceList()
}

func (p *Person) Email() {
	fmt.Printf(emailFormat, p.Name, p.Age)
}

func (p *Person) Older() {
	p.Age++
}
