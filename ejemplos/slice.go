package main

import "fmt"

func main() {
	fmt.Println("Clase Slice")
	numeros := []int{1, 2, 3, 4} // arreglo

	//agregamos alementos con la siguiente funcion
	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)
	numeros = append(numeros, 8)
	numeros = append(numeros, 9)
	numeros = append(numeros, 10)

	//generamos un nuevo slice
	nuevoSlice := numeros[0:5]

	// si modificamos un elemento del slice
	//se vera reflejado en los dos ya que un slice
	// es la referencia a una porcion de un arreglo
	// debe de existir un arreglo para optener la referencia
	numeros[0] = 100
	fmt.Println(numeros)
	fmt.Println(nuevoSlice)
}
