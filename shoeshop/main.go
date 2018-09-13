package shoeshop

import (
	"fmt"
	"os"
	"sort"
	"text/template"
)

const shoeTemplate = `
========== PRICE LIST of {{.Name}} ==========
{{range .Products}}{{leftPad .Model 20}} {{leftPad .Brand 10}} {{.Price |printf "%10.2f"}}
{{end}}=============================================
`

type Shop struct {
	Name     string
	Products map[string]Shoe
	Address
}

type Address struct {
	City   string
	Street string
}

type Shoe struct {
	Id     string
	Model  string
	Brand  string
	Colors []string
	Price  float32
}

func (s *Shop) Add(model string, brand string, price float32, color ...string) {
	if s.Products == nil {
		s.Products = map[string]Shoe{}
	}

	id := model + brand

	s.Products[id] = Shoe{
		Id:     id,
		Model:  model,
		Brand:  brand,
		Colors: color,
		Price:  price,
	}
}

func (s *Shop) DeleteById(id string) {
	fmt.Printf("========== Delete: %v ==========\n", id)
	delete(s.Products, id)
}

func (s *Shop) PriceList() {
	funcs := template.FuncMap{
		"rightPad": rightPad,
		"leftPad":  leftPad,
	}

	t, err := template.New("Shoe").Funcs(funcs).Parse(shoeTemplate)

	if err != nil {
		fmt.Println("Tempalte error: ", err)
	}

	t.Execute(os.Stdout, s)
}

func (s *Shop) FindByBrand(brand string) []Shoe {
	fmt.Printf("========== Find: %v ==========\n", brand)

	var ret []Shoe

	for _, p := range s.Products {
		if p.Brand != brand {
			continue
		}

		ret = append(ret, p)
	}

	return ret
}

func (s *Shop) Brands() []string {
	var ret []string

	for _, p := range s.Products {
		if contains(ret, p.Brand) {
			continue
		}

		ret = append(ret, p.Brand)
	}

	sort.Strings(ret)

	return ret
}

// Better one
func (s *Shop) BrandsAlt() []string {
	brands := map[string]bool{}

	for _, v := range s.Products {
		if !brands[v.Brand] {
			brands[v.Brand] = true
		}
	}

	list := make([]string, len(brands))

	for b := range brands {
		list = append(list, b)
	}

	sort.Strings(list)

	return list
}

func PrintList(s []Shoe) {
	fmt.Printf("%20s %10s %s\n", "MODEL", "BRAND", "COLORS")

	for _, p := range s {
		fmt.Printf("%20s %10s %v\n", p.Model, p.Brand, p.Colors)
	}
}

func contains(h []string, n string) bool {
	for _, a := range h {
		if a == n {
			return true
		}
	}

	return false
}

func rightPad(s string, l int) string {
	return fmt.Sprintf("%[2]*.[2]*[1]s", s, l)
}

func leftPad(s string, l int) string {
	return fmt.Sprintf("%-[2]*.[2]*[1]s", s, l)
}
