package arreglos_slices

import "fmt"

func Capacidad() {
	arreglo1 := make([]int, 10, 15)
	fmt.Printf("Tamaño $d, capacidad $d", len(arreglo1), cap(arreglo1))
}
