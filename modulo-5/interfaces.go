package main

import "fmt"

// interfaz (Que hacer pero no como)
type Animal interface {
	Comer()
	Dormir()
}

// clase perro
type Perro struct {
	Nombre string
}

// se implementa los metodos de la interfaz perro para los metodos
func (self Perro) Comer() {
	fmt.Println("El perro come!")
}

func (self Perro) Dormir() {
	fmt.Println("El perro duerme!")
}

func ejecutarAcciones(animal Animal) {
	animal.Comer()
	animal.Dormir()
}

func main() {
	fmt.Println("Interfaces")

	loki := Perro{Nombre: "Loki"}

	ejecutarAcciones(loki)
}
