package main

import (
	"fmt"
	"os"
)

/*
Ejercicio 2 - Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas
	liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo TXT.
Desarrollar el código necesario para leer los datos de un archivo llamado “customers.txt”.
Sin embargo, debemos tener en cuenta que la empresa no nos ha pasado el archivo a leer por
	el programa. Dado que no contamos con el archivo necesario, se obtendrá un error.
	En tal caso, el programa deberá arrojar un panic al intentar leer un archivo que no existe,
	mostrando el mensaje “El archivo indicado no fue encontrado o está dañado”.
Más allá de eso, deberá siempre imprimirse por consola “Ejecución finalizada”.
*/

func main(){

	


	fmt.Println("\nIniciando...")
	readFile("\ncustomer.txt")
	fmt.Println("\nEjecución finalizada")

}


func readFile(name string){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.Open(name)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
}

//exit status 2 ->siempre devuelve esto, está bien
//panic corta flujo de ejecución del programa, por lo tanto, el defer-recover, permite manejar el panic para que se corte la ejecución. 
//el defer Se define denro del scope de la función donde se va a ejecutar el panic