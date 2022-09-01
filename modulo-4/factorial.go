package main

import "fmt"

func factorial(numero int) int {

	if numero == 0 {
		return 1
	}

	return factorial(numero-1) * numero
}

func main() {
	fmt.Println("Recursividad -- Factorial de un numero entero")

	resultado := factorial(5)

	fmt.Println(resultado)

}
