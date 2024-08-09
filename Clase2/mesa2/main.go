package main

// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
// Categoría C: su salario es de $1.000 por hora.
// Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
// Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

const(
	categoriaA = "categoriaA"
	categoriaB =  "categoriaB"
	categoriaC = "categoriaC"
)

func main(){

}

func salarioA(){}
func salarioB(){}
func salarioC(){}



func calcularSalario(cantHoras int, categoria string){
	switch calculo {
		case categoriaA
		return ((cantHoras * 60) * salarioA())
		case categoriaB
		return ((cantHoras * 60) * salarioB())
		case categoriaC
		return ((cantHoras * 60) * salarioC())
	}
}