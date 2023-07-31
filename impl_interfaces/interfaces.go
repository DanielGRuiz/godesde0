package impl_interfaces

import (
	"fmt"

	"github.com/DanielGRuiz/godesdeo/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("Soy un/a %s y estoy respirando \n", humano.Sexo())
}
