package main

import "fmt"

var username string // vacia

func fun1() {
	username = "Cambie en la funcion 1"
	fmt.Println("Funcion1:", username)
}

func fun2() {
	username = "Cambie en la funcion 2 no se que valor tenndre"
	fmt.Println("Funcion2:", username)
}

func main() {
	fmt.Println("Variables globales")

	username = "Pedrin"

	fun1()
	fun2()
	//Se imprime despues de cambiar su valor en dos ocaciones

	fmt.Println("La variable termino con el valor: ", username)
}
