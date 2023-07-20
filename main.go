package main

import (
	"fmt"
	"runtime"

	"github.com/DanielGRuiz/godesdeo/variables"
)

func main() {
	estado, texto := variables.ConvierteATexto(1577)

	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("No es Windows")
	} else {

		fmt.Println("Es Windows")
	}
}
