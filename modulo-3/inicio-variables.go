package main

import "fmt"

func main() {

	if nombre, edad := "Pedro", 7; nombre == "Pedro" {
		fmt.Println("Hola", nombre, "te damos la bienvenida al curso de Go!")
	} else {
		println("Los valores son: ", nombre, edad)
	}
}
