package main

import "fmt"

func main() {
	meses := []string{"Enero", "Febrero", "Marzo", "Abril",
		"Mayo", "Junio", "Julio", "Agosto", "Septiembre"}

	//capacidad y longitud
	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v, la capacidad es: %v %p\n", longitud, capacidad, meses)

	meses = append(meses, "Octubre") //Si la estructura esta en su maxima capacidad
	//se genera un nuevo arreglo
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")

	// se vuelve a asignar la longitud y la capacidad
	longitud = len(meses)
	capacidad = cap(meses)
	fmt.Printf("La longitud es: %v, la capacidad es: %v %p\n", longitud, capacidad, meses)
}
