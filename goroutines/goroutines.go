package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(3000 * time.Microsecond)
		fmt.Println(letra)
	}

	canal1 <- true
}
