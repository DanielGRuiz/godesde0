package defec_panic

import (
	"fmt"
	"log"
)

func VemosEfec() {
	fmt.Println("Mensaje inicial")

	defer fmt.Println("Mensaje final")

	fmt.Println("Mensaje después del inicial")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Un error que generó Panic\n%v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Prueba con valos 1")
	}
}
