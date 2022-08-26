package main

import "fmt"

func main() {
	fmt.Println("Funciones anonimas")

	func() {
		fmt.Println("Hola mundo desde una funcion sin nombre")
	}()

	miFuncion := func() {
		fmt.Println("Agregando una funcion a una variable")
	}

	miFuncion()
}
