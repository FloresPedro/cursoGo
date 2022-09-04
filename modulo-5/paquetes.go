package main

import (
	"fmt"

	"examplo/CodigoFacilito"
)

func main() {
	fmt.Println("Paquetes en go")
	//folder o carpeta que alacena archivos .go

	curso := CodigoFacilito.Curso{Titulo: "Curso profecinal de Go"}

	fmt.Println(curso.GetTitulo())
	//Si se escribe con mayuscula es publica
	//si se escribe con minuscula es privada
}
