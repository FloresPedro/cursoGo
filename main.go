package main

import "fmt"

func main() {
	fmt.Println("Canales para la concurrencia")

	channel := make(chan string)

	/*go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)*/
	go func() {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}()

	msg := <-channel

	fmt.Println(msg)
}
