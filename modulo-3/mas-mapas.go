package main

func main() {
	println("Obtener valores de un mapa")

	usuarios := make(map[string]string)

	usuarios["Pedro"] = "florescortespedro94@gmail.com"
	//un mapa devuelve dos variables, el valor si existe
	//y otra variable de tipo boolean la cual dice si cierta llave
	//existe con valor, de lo aprendido se puede
	//declarar en el if y validar en cualquiera de sus partes
	if correo, ok := usuarios["Pedro"]; ok {

		println(correo)

	} else {

		println("No fue posible obtener el valor")
	}
}
