package main

import "fmt"

func main() {
	var a int
	a = 10

	//con & antes del nombre de la variable: &a, nos imprime el espacio en memoria de la misma
	fmt.Println("Valor de a es: ", a, "el espacio en memoria: ", &a)


	//en este caso no modifica el valor de a, porque no tiene las referencias del espacio en memoria de la variable a
	//solo tiene una copia de la variable a
	calc(a)
	fmt.Println("Cambio de valor de a a: ", a)

	fmt.Println()
	fmt.Println()

	calc1(&a)
	fmt.Println("Cambio de valor de a a: ", a)


	//declaro puntero
	var ptr *int
	ptr = &a //este es otro puntero que tiene la referencia a la variable a
	fmt.Println(ptr)
	*ptr = 30
	fmt.Println("\nPuntero ptr: ", ptr, "el valor desreferenciado: ", *ptr)
	fmt.Println("\nImprimo a final", a)

}

func calc(a int) {
	a = 20
}

func calc1(a *int) {
	fmt.Println("Lo que me llega: ", a, " y el valor desreferenciado: ", *a)
	*a = 20 //aquí desreferencio espacio en memoria de a
	fmt.Println("Valor de a: ", a, " y el valor desreferenciado: ", *a)
}