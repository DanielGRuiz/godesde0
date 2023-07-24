package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresoNumeros() {
	numero1 := EntradaNumeroPorTeclado("Ingrese un número: ")
	fmt.Println(numero1)

	texto1 := EntradaTextoPorTeclado("Ingrese un texto: ")
	fmt.Println(texto1)
}

func EntradaNumeroPorTeclado(aviso string) int {
	var numero int
	var err error

	fmt.Print(aviso)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Dato ingresado es incorrecto : " + err.Error())
		}
	}

	return numero
}

func EntradaTextoPorTeclado(aviso string) string {
	var texto string

	fmt.Print(aviso)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto = scanner.Text()
	}

	return texto
}
