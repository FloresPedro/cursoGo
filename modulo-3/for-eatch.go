package main

import "fmt"

func main() {
	animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}
	//for tradicional
	/*
		for indice := 0; indice < len(animales); indice++ {

			elemento := animales[indice]

			fmt.Println(elemento)
		}*/

	//for eatch
	// en dado caso de que no se quiera utilizar una
	// variable solo se pone un _ en  su lugar
	for indice, elemento := range animales {
		fmt.Println("El valor es:", indice, "El valor es: ", elemento)
	}
}
