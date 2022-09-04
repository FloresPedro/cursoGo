package main

import "fmt"

type Animal interface {
	Dormir()
}

type Mascota interface {
	Comer()
}

type Gato struct {
	Nombre string
}

func (this Gato) Dormir() {
	fmt.Println("El animal se durmio")
}

func (this Gato) Comer() {
	fmt.Println("El animal Come")
}

func funcionParaUnaMascota(animal Mascota) {

	fmt.Println("El objeto es una mascota")
}

func funcionParaUnAnimal(animal Animal) {
	fmt.Println("El objeto es un animal")
}

func main() {
	fmt.Println("Multiples interfaces")
	gato := Gato{Nombre: "Mi gato"}
	funcionParaUnAnimal(gato)
	funcionParaUnaMascota(gato)
	gato.Dormir()
	gato.Comer()

}
