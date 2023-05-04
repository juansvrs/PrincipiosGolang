package main
import "fmt"
func main(){
	saludo:="hola"
	saludo2:="hola amigo!"

	//println: imprime en una linea aparte
	//print imprime en el siguiente espacio
	fmt.Println(saludo)
	fmt.Print(saludo2)

	nombre:="Juan"
	edad:=22

	fmt.Printf("El es %s y su edad es %d \n",nombre,edad)

	mensaje := fmt.Sprintf("El es %s y su edad es %d ",nombre,edad)

	fmt.Println(mensaje)

	//conocer el tipo de dato de una variable:
	fmt.Println("nombre: %T",nombre)

	//ingresar datos por teclado
	fmt.Println("ingrese otro nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("El nombre digitado es: ",nombre)
}