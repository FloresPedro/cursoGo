package main

import "fmt"

func main() {

	fmt.Println("Clase de slice con la funcion make")
	//primer argumento tipo de slice
	//segundo argumento longitud del slice
	//tercer argumento capacidad maxima
	slice := make([]int, 3, 3)

	slice[0] = 100
	slice[1] = 200
	slice[2] = 300

	//al agregar un nuevo elemento y superar la
	//capacidad maxima se crea un nuevo slice con el doble
	//de la capacidad maxima
	slice = append(slice, 400)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
