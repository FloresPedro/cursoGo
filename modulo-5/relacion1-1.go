package main

import "fmt"

type User struct {
	Name   string
	Emain  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

func main() {
	fmt.Println("Estructuras enbebidas")

	eduardo := User{Name: "Eduardo", Emain: "eduardo@cody.com", Active: true}

	pedro := User{Name: "Pedro", Emain: "pedrin@cody.com", Active: true}

	pedroStudent := Student{User: pedro, Id: "1f1"}

	fmt.Println(eduardo)

	fmt.Println(pedroStudent)

}
