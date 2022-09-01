package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (this *User) setName(name string) {
	this.Name = name
}

func (this *User) getName() string {
	return this.Name
}

func main() {
	fmt.Println("Metodos en go")

	defer func() { //funcion que se ejecuta al final de metodo
		fmt.Println("Mi...voy")
	}()

	cody := User{Name: "Pedrin", Email: "info@codigofacilito.com"}

	cody.setName("Sofia")

	fmt.Println(cody.getName())

}
