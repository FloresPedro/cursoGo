package main

import "fmt"

func main() {

	numero := 12345

	contador := 0
	//Go soporta la palabra while
	for numero > 0 {
		numero = numero / 10

		contador++
	}

	fmt.Println("La cantidad de digitos es: ", contador)
}
