package main

import "fmt"

func main() {

	fmt.Println("Mapas en Go")
	//mapa de enteros que guardan los dias
	dias := make(map[int]string)

	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	//imprime los dias
	fmt.Println(dias)
	//eliminar una llave con su respectivo valor
	delete(dias, 4)
	//reimprimir los dias
	fmt.Println(dias)

	//otro ejemplo de mapas
	//map: palabra reservada,
	usuarios := make(map[string][]int)
	usuarios["Pedro"] = []int{10, 9, 8, 10, 5}

	fmt.Println(usuarios)
}
