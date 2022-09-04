package main

import "fmt"

type Curso struct { // un curso puede tener varios videos
	Titulo string
	Videos []Video
}

type Video struct { // un video le pertenece a un curso
	Titulo string
	Curso  Curso
}

func main() {
	fmt.Println("Estructuras enbebidas relacion uno-muchos")
	//Se declaran los videos
	video1 := Video{Titulo: "1.- Introduccion"}
	video2 := Video{Titulo: "2.- Intalacion"}
	//se crea el arreglo de videos
	videos := []Video{video1, video2}
	//se establece la relacion
	curso := Curso{Titulo: "Curso profecional de Go", Videos: videos}
	//se asigna el curso a los videos
	video1.Curso = curso
	video2.Curso = curso
	// se imprime la relacion
	fmt.Println(video1)
	fmt.Println(video2)
	fmt.Println(curso)
	// se imprime el titulo del curso
	fmt.Println(video1.Curso.Titulo)
	//se imprumen los titulos de los videos
	for _, video := range curso.Videos {
		fmt.Println(video.Titulo)
	}
}
