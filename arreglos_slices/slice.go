package arreglos_slices

import "fmt"

func Capacidad() {
	arreglo1 := make([]int, 10, 15)
	fmt.Printf("arreglo1 => Tamaño %d, capacidad %d", len(arreglo1), cap(arreglo1))

	numeros := make([]int, 0, 0)
	fmt.Printf("numero => Tamaño %d, capacidad %d", len(numeros), cap(numeros))

	for i := 0; i < 10000000; i = 10000 {
		numeros = append(numeros, i)
		fmt.Printf("\n numero => Tamaño %d, capacidad %d", len(numeros), cap(numeros))
	}
}
