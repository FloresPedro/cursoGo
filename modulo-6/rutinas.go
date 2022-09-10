package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Go rutines")
	go mensaje_lentoooo("Pedro")

	fmt.Println("Que aburrido esperar esto")
	var wait string
	fmt.Scanln(&wait)
}

func mensaje_lentoooo(nombre string) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
