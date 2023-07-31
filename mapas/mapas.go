package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Colombia"] = "Bogotá"
	paises["Venezuela"] = "Caracas"
	paises["Mexico"] = "Mexico"

	fmt.Println(paises)
	fmt.Println(paises["Colombia"])

	equipos := map[string]int{
		"equipo1": 2,
		"equipo2": 4,
		"equipo3": 8,
		"equipo4": 16,
		"equipo5": 32,
		"equipo6": 64,
		"equipo7": 128,
	}

	fmt.Println(equipos)

	for equipo, puntaje := range equipos {
		fmt.Printf("-> %s, puntaje %d\n", equipo, puntaje)
	}

	delete(equipos, "equipo5")
	fmt.Println(equipos)

	equipo, existe := equipos["equipo2"]
	fmt.Println(equipo, existe)
}
