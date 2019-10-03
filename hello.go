package main

import (
	"fmt"

	"github.com/manuelmunoz00/hello/numbers"

	"github.com/manuelmunoz00/hello/stringutil"
)

func main() {
	fmt.Println("Hello")
	github := "buhtig"
	fmt.Println(stringutil.Reverse(github))
	fmt.Println(stringutil.Mayusculas(github))
	fmt.Println(numbers.EsMayorEdad(20))
}
