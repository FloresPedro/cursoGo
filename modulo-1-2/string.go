package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("String clase cadenas")
	//declarar una variable String
	var curso1 string = "Curso profecional de Go!"
	//variable que el compilador sepa el tipo de dato
	var curso2 = "Curso de Go que el compilador sepa el tipo de dato"
	//asignacion curso
	curso3 := "Curso profecional de Go!"

	longitud := len(curso1)

	fmt.Println(curso1, longitud)
	fmt.Println(curso2)
	fmt.Println(curso3)

	fmt.Println("Se optiene el ultimo caracter")
	primerCaracter := curso1[0] //Extraes el primer elemento de la lista
	ultimoCaracter := curso1[len(curso1)-1]

	fmt.Println(primerCaracter) //devuelve 67 por sel en uint8

	fmt.Println(reflect.TypeOf(primerCaracter))
	fmt.Printf("%c\n", primerCaracter) // aqui si se muestra el valor de la primera letra

	fmt.Printf("%c\n", ultimoCaracter) // imprimimos el ultimo caractr
}
