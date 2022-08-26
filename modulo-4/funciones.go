package main

import "fmt"

func saludar(username string) {

	fmt.Println("Hola ", username)
}

func main() {
	fmt.Println("Funciones...")

	saludar("Cody")
}
