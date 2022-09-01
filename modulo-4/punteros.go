package main

import "fmt"

func modificarValor(parametro *string) {
	*parametro = "Cambio de valor"
}

func main() {

	fmt.Println("Punteros")

	var curso = "Curso profecional de Go!"

	fmt.Println("Antes de la funcion", curso)

	modificarValor(&curso) // pasamos la referencia al espacio en memoria

	fmt.Println("Despues de la funcion: ", curso)
}
