package main

import "fmt"

func main() {
	fmt.Println("Funcion panic")

	var dividendo, divisor int

	fmt.Println("Ingresa un valor para el dividendo: ")

	fmt.Scanf("%d", &dividendo)

	fmt.Println("Ingresa un valor para el dividendo: ")
	fmt.Scanf("%d", &divisor)

	if divisor == 0 {
		panic("Se intento dividir enter cero")
	}

	resultado := dividendo / divisor

	fmt.Println(resultado)
}
