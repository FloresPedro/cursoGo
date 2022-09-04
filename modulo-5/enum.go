package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	UserName string
	Type     UserType
}

func main() {

	fmt.Println("Enums")
	pedrin := User{UserName: "Pedrin", Type: Student}
	sofia := User{UserName: "Sofia", Type: Teacher}

	fmt.Println(pedrin)
	fmt.Println(sofia)

	if pedrin.Type == Student {
		fmt.Println("El usuario", pedrin.UserName, "es un estudiante")
	}
}
