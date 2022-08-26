package main

import "fmt"

func main() {

	var numero = 1

	for {
		fmt.Println("Ciclo infinito: ", numero)
		numero++
		if numero == 1000000 {
			break
		}
	}
}
