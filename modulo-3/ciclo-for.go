package main

import "fmt"

func main() {
	fmt.Println("ciclo for...")

	for i := 0; i <= 100; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
