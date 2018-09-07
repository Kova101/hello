package main

import (
	"fmt"
	"os"
)

const name = "boborján"
const num = 42
const other = 3.14
const emailTemplate = `
DEAR $USER

Please send all yur money to me!!!

best regards,
Broli
`

func main() {
	fmt.Printf("%v, %v, %v\n", name, num, other)
	fmt.Printf("%T, %T, %T\n", name, num, other)
	fmt.Println(os.ExpandEnv(emailTemplate))
}
