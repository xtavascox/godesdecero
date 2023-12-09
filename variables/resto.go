package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestVariables() {
	Nombre = "Gustavo"
	Estado = true
	Sueldo = 5000.50
	Fecha = time.Now()

	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Estado: ", Estado)
	fmt.Println("Sueldo: ", Sueldo)
	fmt.Println("Fecha: ", Fecha)
}

func TextConvert(numero int) (bool, string) {
	var text string
	text = strconv.Itoa(numero)
	return true, text
}
