package main

import "fmt"

func factorial(numero int) int {

	if numero == 0 {
		return 1
	}

	return factorial(numero-1) * numero
}

// tipo de dato personalizado
type customeFunction func(n int) int

func main() {

	//declaramos una variable de mi tipo de dato
	var miFuncion customeFunction
	//evalumos si es igual a nulo<nil>
	if miFuncion == nil {
		//asignamos algo a la variable
		//en este caso una funcion factorial
		miFuncion = factorial
	}
	//la ejecutamos si requiere argumentos
	resultado := miFuncion(5)

	fmt.Println(resultado)
}
