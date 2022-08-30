package main

import "fmt"

//Funciones que retornan valores
//esta medio raro en como lo hace
//si vienes de java o de algun otro lenguaje

// func suma(numero1 int, numero2 int) int {
// si los argumentos son del mismo tipo de dato se puede poner
// solo una vez al final del ultimo algumento
// las funciones pueden regresar mas de un argumento
// que loco
func suma(numero1, numero2 int) (int, string) {
	return numero1 + numero2, "Valor verificado"
}

func main() {
	// se extraen los dos resultados de la funcion
	resultado, ok := suma(10, 20)
	//resultado, _ := suma(15, 20) para NO hacer uso del segundo argumento

	fmt.Println(resultado, " string de la funcion ", ok)
}
