package ejercicios

import "strconv"

func GreaterOrLess(value string) (int, string) {
	numberTransform, err := strconv.Atoi(value)
	if err == nil {
		if numberTransform > 100 {
			return numberTransform, "El valor es mayor a 100"
		} else if numberTransform < 100 {
			return numberTransform, "El valor es menor a 100"
		} else {
			return numberTransform, "El valor es igual a 100"
		}
	}
	return numberTransform, "El valor no es un número"
}
