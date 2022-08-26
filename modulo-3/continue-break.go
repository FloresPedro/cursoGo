package main

import "fmt"

func main() {
	// break -- continue

	for i := 1; i < 10; i++ {

		if i == 5 {
			fmt.Println("Se continuo")
			continue
		}

		if i == 8 {
			fmt.Println("Cieta condicion termino el ciclo")
			break
		}
		fmt.Println(i)
	}
}
