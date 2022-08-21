package main

import "fmt"

/*
&& and
|| or
! negacion
*/

func main() {
	fmt.Println("Operadores logicos 2")
	resultado := 20 == 20 && 30 == 30
	resultado2 := 50 != 50 && 40 == 40 || (15 < 20 && 9 == 8)
	fmt.Println(resultado)
	fmt.Println(resultado2)
}
