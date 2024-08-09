package main

import "fmt"

func main() {

	ejercicio2()


}

func ejercicio1(){
	var palabra string = "palabra"
	//palabra2 := "palabra2"

	//cant := len(palabra)
	var cant int = len(palabra)

	fmt.Println("La palabra" , palabra , "tiene" , cant , "letras")
	fmt.Printf("La palabra %s tiene %d letras", palabra, cant)
	fmt.Printf("")

	for i := 0; i < cant; i++ {
		fmt.Println(string(palabra[i]))
	}
}

func ejercicio2(){
	var meses = [12]string{"Enero", "febrero", "marzo", "Enero", "febrero", "marzo", "Enero", "febrero", "marzo", "Enero", "febrero", "marzo"}
	//slice es cuando no declaramos la cantidad, corchetes vacios, sino es un arrays

	mes := 0
	if mes <1 || mes >12{
		fmt.Println("Error: fuera de mes")
	}


		if mes >= 1 && mes <=12 {

		for i := 0; i <= len(meses); i++{
			if i == mes{
				fmt.Println(meses[i-1])
			}
		}
	}
if mes >= 1 && mes <=12 {

	fmt.Println(meses[mes-1])
	}	
}