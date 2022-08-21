package main

import "fmt"

/*
> mayor que, >= mayor o igual
< menor que, <= menor o igual
== igual, != diferente, !neegacion?
*/
//Se pueden comparar cadenas de caracteres!!
func main() {
	fmt.Println("Operadores Relacionales ")
	var edad = 10
	var resultado = edad < 5 //Bool
	fmt.Println(!resultado)
}
