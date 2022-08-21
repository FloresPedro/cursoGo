package main

import "fmt"

func main() {
	monedas := [...]string{0: "Dolar canadiense", 1: "Peso mexicano",
		2: "Dolar", 5: "Euro"}
	fmt.Println(monedas)
	//crear un slice
	subMonedas := monedas[0:2]
	fmt.Println(subMonedas)
}
