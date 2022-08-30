package main

import "fmt"

// tipos de datos (Nuevo)
// cualquier funcion que tenga dos argumentos de tipo entero
// y retorne un entero sera considerada una operacion
type Operacion func(balance, cantidad int) int

func suma(numero1, numero2 int) int { // es de tipo operacion
	return numero1 + numero2
}

func resta(numero1, numero2 int) int { // es de tipo operacion
	return numero1 - numero2
}

func ejecutarOperacion(funcion Operacion, balance, cantidad int) {
	fmt.Println("Antes de la operacion")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es:", resultado)

	fmt.Println("Despues de la operacion")
}

func main() {
	fmt.Printf("Funciones como argumentos")
	//introduccion a programacion funcional en Golang
	ejecutarOperacion(resta, 10, 50)
}
