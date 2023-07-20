package ejercicios

import (
	"strconv"
)

func DevuelveValores(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	var mensaje string

	if err != nil {
		return 0, "Hubo un error"
	}

	if numero > 100 {
		mensaje = "Es mayor a 100"
	} else {
		mensaje = "Es menor a 100"
	}

	return numero, mensaje
}
