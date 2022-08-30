package main

import "fmt"

func modificarVariable(parametro string) {
	fmt.Println("Valor actual:", parametro)
	parametro = "Nuevo curso de Go!"
	fmt.Println("Nuevo valor", parametro)

	fmt.Printf("%p\n", &parametro)
}

func main() {
	fmt.Println("Variables y scope")

	var curso = "Curso profecional de Go!"

	modificarVariable(curso)

	fmt.Println(">>>", curso)

	fmt.Printf("%p\n", &curso)

}
