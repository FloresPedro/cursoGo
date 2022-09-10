package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func crearUsuario() {

	clearConsole()

	fmt.Print("Ingresa un valor para username: ")
	username := readLine()

	fmt.Print("Ingresa un valor para email: ")
	email := readLine()

	fmt.Print("Ingresa un valor para age: ")
	//equivalente a un Parse int
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No se pudo convertir la edad")
	}

	id++

	user := User{id, username, email, age}

	users[id] = user

	fmt.Println(">>> Se creara un nuevo usuario... \n")
}

func listarUsuarios() {
	clearConsole()
	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}
	fmt.Println("\n")
}

func actualizarUsuario() {

	fmt.Println("Se actualizara el usuario..")
}

func eliminarUsuario() {
	clearConsole()
	fmt.Print("Ingresa el id del usuario a eliminar: ")
	if id, err := strconv.Atoi(readLine()); err != nil {
		panic("error al parsear el numero")
	} else if _, ok := users[id]; ok {
		delete(users, id)
	}
	fmt.Println(">>> Se elimino el usuario de forma correcta...\n")
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {
	//lee de consola lo que ingreses asta que preciones un salto de linea
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor!")
	} else {
		// retorna la cadena ingresada sin el salto de linea
		return strings.TrimSuffix(option, "\n")
	}
}

func main() {
	fmt.Println("Proyecto final")

	var option string

	users = make(map[int]User)
	//para lectura de datos io
	reader = bufio.NewReader(os.Stdin)
	for {

		fmt.Println("A) Crar")

		fmt.Println("B) Listar")

		fmt.Println("C) Actualizar")

		fmt.Println("D) Eliminar")

		fmt.Print("Ingresa una opcion: ")

		option = readLine()

		if option == "quit" || option == "q" || option == "c" {
			break
		}

		switch option {
		case "a", "crear", "A":
			crearUsuario()
		case "b", "listar", "B":
			listarUsuarios()
		case "c", "actualizar", "C":
			actualizarUsuario()
		case "d", "eliminar", "D":
			eliminarUsuario()
		default:
			fmt.Println("Selecciona una opcion correcta por favor")
		}
	}

	fmt.Println("Adios")

}
