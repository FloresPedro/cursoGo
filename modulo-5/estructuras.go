package main

import "fmt"

// para crear nuestras propias estructuras
// Es un tipo de dato
type User struct {
	Name  string // se inicializan con un string vacio
	Email string
}

func main() {
	fmt.Println("Manejo de estructuras")

	var cody User

	cody.Name = "Pedrin"
	cody.Email = "pedrin@gmail.com"

	fmt.Println(cody)
	fmt.Println(cody.Name)
	fmt.Println(cody.Email)

	fmt.Println("Otra forma de definir un objeto")
	//segunda forma de definir un objeto
	usuario2 := User{Name: "Cody", Email: "cody@codigofacilito.com"}

	fmt.Println(usuario2)
	//tercera forma de crear un objeto
	Name := "Sofia"
	Email := "sofi@gmail.com"

	usuario3 := User{Name, Email}

	fmt.Println(usuario3)
}
