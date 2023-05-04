package main

import "fmt"

func main() {
	/* GO nos obliga a utilizar las variables
	porque sino, ocupamos memoria, por eso sale error*/
	var nombre1 string
	nombre1 = "Juan Luis"

	var nombre2 string = "Roel"

	/* simplificacion de la sintaxis*/
	edad := 22

	fmt.Println(nombre1, nombre2, edad)

	/*mas tipos de datos, definiendo*/
	/* cuando se definen asi, a diferencia de otros lenguajes
	si empiezan con un valor inicial y no un NULL o vacio*/
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a,b,c,d)

const pi =3.141592 /*no se modifica su valor*/

}
