package main

import (
	"github.com/DanielGRuiz/godesdeo/webserver"
)

func main() {
	/*
		estado, texto := variables.ConvierteATexto(1577)

		fmt.Println(estado)
		fmt.Println(texto)

		if os := runtime.GOOS; os == "windows" {
			fmt.Println("Es Windows")
		} else {
			fmt.Println("No es Windows")
		}

		valor, mensaje := ejercicios.DevuelveValores("asd")
		fmt.Println(valor)
		fmt.Println(mensaje)

		teclado.IngresoNumeros()
	*/

	//fmt.Printf(ejercicios.TablaMultiplicar())

	//files.CrearArchivo()

	//files.AgregarArchivo()

	//files.LeerArchivo()

	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.AgregarUsuario()

	//daniel := new(modelos.Hombre)
	//impl_interfaces.HumanosRespirando(daniel)

	//liliana := new(modelos.Mujer)
	//impl_interfaces.HumanosRespirando(liliana)

	//defec_panic.EjemploPanic()

	//canal1 := make(chan bool)
	//go goroutines.MiNombreLento("Daniel Ruiz", canal1)

	//defer func() {
	//	<-canal1
	//}()
	//fmt.Println("Esperando ...")
	//var valorEntrada string
	//fmt.Scan(&valorEntrada)

	webserver.MiWebServer()
}
