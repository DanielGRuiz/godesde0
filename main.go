package main

import (
	"fmt"
	"runtime"

	"github.com/DanielGRuiz/godesdeo/ejercicios"
	"github.com/DanielGRuiz/godesdeo/variables"
)

func main() {
	estado, texto := variables.ConvierteATexto(1577)

	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "windows" {
		fmt.Println("Es Windows")
	} else {
		fmt.Println("No es Windows")
	}

	valor, mensaje := ejercicios.DevuelveValores("asd")
	fmt.Println(valor)
	fmt.Println(mensaje)

	IngresoNumeros()
}
