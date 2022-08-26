package main

import "fmt"

func main() {

	fmt.Println("Condicionales -- Controles de flujo")

	var edad int

	println("Ingresa tu edad: ")

	fmt.Scanf("%d", &edad)

	if edad >= 18 {
		println("Eres un adulto :-)")
	} else {
		println("Aun no cumples la mayoria de edad :v")
	}
}
