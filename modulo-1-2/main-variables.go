package main

import "fmt"

//go build main.go -> Para archivo ejecutable
//go run main.go -> para ejecutar el programa sin generar el compilado

func main() {
	/*
		//declarar variables
		var nombre string
		var edad int

		//asignar valores a las variables
		nombre = "Pedro"
		edad = 28
	*/

	//declarar variable que el compilador intuya el tipo
	nombre := "Pedro"
	edad := 26

	var altura = 1.83 //float
	fmt.Println("El nombre es:", nombre, "la edad es:", edad, "La altura es:", altura)

}
