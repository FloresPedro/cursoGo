package main

import "fmt"

func main() {

	var calificacion int

	println("Ingresa una calificacion: ")

	fmt.Scanf("%d", &calificacion)
	//estructura de un switch case
	println("Primer switch.....")
	switch {
	case calificacion == 10:
		println("Felicidades, obtiviste una calificacion perfecta")
	case calificacion == 9 || calificacion == 8:
		println("Aprobaste la materia pero se te escapo la gloria")
	case calificacion == 6 || calificacion == 7:
		println("Casi no pasas la materia pero eres todo un soldado")
	case calificacion >= 0 && calificacion <= 5:
		println("No aprovaste la materia")
	default:
		println("Calificaicon no valida")
	}

	/*if calificacion == 10 {
		println("Felicidades, obtiviste una calificacion perfecta")
	} else if calificacion == 9 || calificacion == 8 {
		println("Aprobaste la materia pero se te escapo la gloria")
	} else if calificacion == 6 || calificacion == 7 {
		println("Casi no pasas la materia pero eres todo un soldado")
	} else if calificacion >= 0 && calificacion <= 5 {
		println("No aprovaste la materia")
	} else {
		println("Calificaicon no valida")
	}*/

	//otra forma de usar un switch
	println("Segundo switch.....")
	switch calificacion {
	case 10:
		println("Felicidades, obtiviste una calificacion perfecta")
	case 9, 8:
		println("Aprobaste la materia pero se te escapo la gloria")
	case 6, 7:
		println("Casi no pasas la materia pero eres todo un soldado")
	case 0, 1, 2, 3, 4, 5:
		println("No aprovaste la materia")
	default:
		println("Calificaicon no valida")
	}
}
