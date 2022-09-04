package main

import "fmt"

func mostrarVariable(objeto interface{}) {
	fmt.Printf("El valor de la variable es: %v\n", objeto)
}

func suma(numero1, numero2 int) int {
	return numero1 + numero2
}

type User struct {
	UserName string
}

func main() {
	fmt.Println("Interfaces bacias")
	//En go todas las estructuras implementan una interfaz vacia

	pedrin := User{UserName: "Pedrin"}
	mostrarVariable(true)
	mostrarVariable(suma)
	mostrarVariable(pedrin)
}
