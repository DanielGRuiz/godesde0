package main

import (
	"fmt"

	"github.com/DanielGRuiz/godesdeo/variables"
)

func main() {
	estado, texto := variables.ConvierteATexto(1577)

	fmt.Println(estado)
	fmt.Println(texto)
}
