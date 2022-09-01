package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups al parecer el programa no finalizo de forma correcta")
		}
	}()

	//Programar funciones
	//y manejor de errores al estilo Go
	//leemos el archivo

	//validamos si existe error
	if file, err := os.Open("examples.txt"); err != nil {
		panic("Ocurrio un error al intentar abrir el archivo")
	} else {
		//funcion que se ejecuta al finalizar el bloque de codigo main
		//para cerrar el archivo abierto
		defer func() {
			fmt.Println("Cerramos el archivo!!")
			file.Close()
		}()
		//Creamos un slice de bytes
		contenido := make([]byte, 254)
		//leemos los bytes en el archivo
		longitud, _ := file.Read(contenido)
		//convertimos a string los bytes
		contenidoArchivo := string(contenido[0:longitud])

		fmt.Println(contenidoArchivo)
	}
	fmt.Println("Tome el contro del programa??")
}
