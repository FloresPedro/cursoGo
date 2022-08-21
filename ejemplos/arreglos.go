package main

import "fmt"

func main() {
	fmt.Println("Arreglos en Go")
	//Se crea un arreglo con 4 elementos enteros
	var numeros [4]int //al inicializar de esta forma todos los valores son 0
	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300
	numeros[3] = 400
	//clear un arreglo definiendo sus valores
	numeros2 := [3]int{1000, 2000, 3000}
	//crear un arreglo sin tener en claro los valores
	numeros3 := [...]int{1, 2, 3}

	fmt.Println(numeros)
	fmt.Println(numeros2)
	fmt.Println(numeros3)
}
