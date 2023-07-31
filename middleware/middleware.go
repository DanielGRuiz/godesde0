package middleware

import "fmt"

func sumar(a, b int) int {
	fmt.Printf("%d + %d = %d | ", a, b, a+b)

	return a + b
}

func restar(a, b int) int {
	fmt.Printf("%d - %d = %d | ", a, b, a-b)

	return a - b
}

func multiplicar(a, b int) int {
	fmt.Printf("%d x %d = %d | ", a, b, a*b)

	return a * b
}

func MiMiddleware() {
	var resultado int
	fmt.Println("Inicio...")

	resultado = operacionesMidd(sumar)(2, 3)
	resultado = operacionesMidd(restar)(2, 3)
	resultado = operacionesMidd(multiplicar)(2, 3)

	fmt.Println("Fin y valor final:", resultado)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")

		return f(a, b)
	}
}
