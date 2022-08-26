package main

import "fmt"

func main() {

	var calificacion int

	println("Ingresa una calificacion: ")

	fmt.Scanf("%d", &calificacion)

	if calificacion == 10 {
		println("Felicidades, obtiviste una calificacion perfecta")
	} else if calificacion == 9 || calificacion == 8 {
		println("Aprobaste la materia pero se te escapo la gloria")
	} else if calificacion == 6 || calificacion == 7 {
		println("Casi no pasas la materia pero eres todo un soldado")
	} else if calificacion >= 0 && calificacion <= 5 {
		println("No aprovaste la materia")
	} else {
		println("Calificaicon no valida")
	}

}
