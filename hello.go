package main

import (
	"fmt"

	"github.com/manuelmunoz00/packagesgo/saludo"

	"github.com/manuelmunoz00/packagesgo/numbers"

	"github.com/manuelmunoz00/packagesgo/stringutil"
)

func main() {
	fmt.Println("Hello")
	github := "buhtig"
	fmt.Println(stringutil.Reverse(github))
	fmt.Println(stringutil.Mayusculas(github))
	fmt.Println(numbers.EsMayorEdad(20))
	fmt.Println(saludo.GetSaludo())
	fmt.Println(saludo.GetVerdad())

}
