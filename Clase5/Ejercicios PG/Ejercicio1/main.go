package main

import "fmt"

/*
Ejercicio 1 - Impuestos de salario
En la función main, definir una variable salary y asignarle un valor de tipo “int”.
Luego, crear un error personalizado con un struct que implemente Error() con el mensaje
“Error: el salario ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que
salary sea menor a 150.000. De lo contrario, imprimir por consola el mensaje
“Debe pagar impuesto”.
*/

//usamos una estructura personalizada para implementar la interfaz Error
type errSalary struct {
	message string
}


//implementamos interfaz Error  travez de nuestra estructura errSalary
func (e errSalary) Error ()string {
	return e.message
}


func main(){

	var salary = 200000
	err := impuesto(salary)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Printf("\n%d Debe pagar impuesto", salary)
	}


}

func impuesto(salary int) error{

	if salary < 150000 {
		return errSalary{message: "\nError: el salario ingresado no alcanza el mínimo imponible."}
	}
	return nil
}

