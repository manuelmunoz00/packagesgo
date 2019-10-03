package numbers

// EsMayorEdad funcion si es mayor de edad
func EsMayorEdad(edad int) bool {
	var retorno bool
	if edad > 18 {
		retorno = true
	} else {
		retorno = false
	}
	return retorno
}
