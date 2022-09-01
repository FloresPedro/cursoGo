package main

import "fmt"

func promedio(calificaciones ...int) int {

	var promedio, suma int

	for _, calificacion := range calificaciones {
		suma = suma + calificacion
	}

	promedio = suma / len(calificaciones)
	return promedio
}
func main() {
	//variadic function
	//recibe n cantidad de argumentos
	fmt.Println("Variadic functions (argumentos variables)")

	resultado := promedio(10, 10, 10)

	fmt.Println("El promedio es:", resultado)
}
