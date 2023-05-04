package main

import "fmt"

func main() {
	a := 20
	b := 10

	//suma
	suma := a + b
	resta := a - b
	div := a / b
	/* para obtener una division exacta se debe hacer typecasting*/
	fmt.Println("suma", suma, " ", "resta", resta, " ", "division", div)

	//modulo
	modulo := a % b
	fmt.Println(modulo)
}
