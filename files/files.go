package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DanielGRuiz/godesdeo/ejercicios"
)

var ruta string = "./files/txt/tabla.txt"

func CrearArchivo() {
	texto := ejercicios.TablaMultiplicar()

	archivo, err := os.Create(ruta)
	if err != nil {
		fmt.Println("Error en creación del archivo:" + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func AgregarArchivo() {
	texto := ejercicios.TablaMultiplicar()

	archivo, err := os.OpenFile(ruta, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error en aperturra del arhivo:" + err.Error())
		return
	}

	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error en la escritura del archivo:" + err.Error())
		return
	}

	archivo.Close()
}

func LeerArchivo() {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error en la lectura del arhivo:" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println("> " + linea)
	}

	archivo.Close()
}
