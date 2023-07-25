package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	for true {
		numero, esNUmero := EntradaNumero("Ingrese un numero:")

		if !esNUmero {
			continue
		}

		for i := 1; i <= 10; i++ {
			resultado := i * numero
			fmt.Println(numero, " x ", i, " = ", resultado)
		}

		break
	}

}

func EntradaNumero(mensaje string) (int, bool) {
	var numero int
	var err error
	var esNumero bool = true

	fmt.Print(mensaje)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Dato ingresado es incorrecto...!")
			esNumero = false
		}
	}

	return numero, esNumero
}
