package main

import "fmt"

func main() {
	fmt.Println("Ingresar informacion por el teclado")
	//declaramos nuestas variables
	var nombre string
	var edad int
	var altura float32
	fmt.Print("Ingresa tu nombre: ")
	//scanf para guardar la info que se ingrese por teclado
	fmt.Scanf("%s", &nombre)
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Print("Ingresa tu altura: ")
	fmt.Scanf("%f", &altura)
	fmt.Printf("Hola %s tu edad es de %d y mides %.2f \n", nombre, edad, altura)

}
