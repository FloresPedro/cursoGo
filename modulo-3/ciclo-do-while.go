package main

import "fmt"

func main() {

	var numero = 10

	//ciclo Do while
	for ok := true; ok; ok = numero < 10 {
		fmt.Println(numero)
		numero++
	}
}
