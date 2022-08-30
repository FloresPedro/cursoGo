package main

import "fmt"

func main() {
	fmt.Println("Funciones anonimas")
	//Funcion anonima y ejecutada
	func() {
		fmt.Println("Hola mundo desde una funcion sin nombre")
	}()
	//funcion anomina declarada en una variable**se puede invocar desde donde sea
	miFuncion := func(username string) {
		fmt.Println("Agregando una funcion a una variable y recibiendo un parametro", username)
	}

	miFuncion("Pedro")

	//funcion que retorna un string
	miFuncionRetorno := func(username string) string {
		message := fmt.Sprintf("Hola %s, te saludamos desde una funcion sin nombre", username)
		return message
	}

	mensajeFinal := miFuncionRetorno("Pedro")

	fmt.Printf(mensajeFinal)

}
