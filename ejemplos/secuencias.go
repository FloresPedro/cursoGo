package main

import "fmt"

const (
	/*
		Domingo   int = 0
		Lunes     int = 1
		Martes    int = 2
		Miercoles int = 3
		Jueves    int = 4
		Viernes   int = 5
		Sabado    int = 6
	*/ // se puede utilizar la palabra
	//reservada iota para expecificar una secuencia
	Domingo int = iota // inicio en 0, iota + n para que sea otro valor
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Secuencias")
	fmt.Println(Domingo)
}
