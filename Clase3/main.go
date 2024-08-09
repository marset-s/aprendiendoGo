package main

import (
	"encoding/json"
	"fmt"
	"log"
	
)

//las estructuras se definen fuera del main
//compocición en Go, permite encapsular estructuras dentro de otras y reutilizar código (similar herencia en java)

//etiquetas: se usa json como etiqueta, y se le indique como devolver los datos en formato json, como tiene que llmarase las claves cuando se parse a json
//con omitempty, si un campo lega vacío, le decimos que lo ignore
// con el guión: "-", le indicamos que no muestre ese valor.. no imprime la fecha

type Person struct {
	Id          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	DateOfBirth string `json:"-"`
}

type Employee struct {
	Id       int    `json:"id"`
	Position string `json:"position"`
	Person   Person `json:"person"`
}


func main(){

	p1 := Person{1, "Juan", "10/10/2000"}
	e1 := Employee{1, "Gerente", p1}

	//llamada el método que retorna string
	fmt.Println(e1.PrintEmployee())

	//llamada al método que imprime simplimente los datos
	e1.PrintEmployee2()


	//
	employee, err := json.Marshal(e1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(string(employee))

	fmt.Println()
	fmt.Println()

	var e2 Employee

	//employee2 := "{\"id\":2,\"position\":\"Developer\",\"person\":{\"id\":2,\"name\":\"Mariano\"}}"
	data := `{"Id":1,"Position":"Gerente","Person":{"Id":1,"Name":"Juan","DateOfBirth":"10/10/2000"}}`

	if err2 := json.Unmarshal([]byte(data), &e2); err2 != nil{
		log.Fatal(err2)
	}

	fmt.Println(e2.PrintEmployee())
}


//método asociado a estructura Employee, que retorna string
func (e Employee) PrintEmployee() string {
	return fmt.Sprintf("Employee ID: %d\n Name: %s\n Date of Birth %s\n Position: %s", e.Id, e.Person.Name, e.Person.DateOfBirth, e.Position)
}

//método asociado a estructura Employee, que no retorna datos sino que imprime los datos
func (e Employee) PrintEmployee2()  {
	fmt.Printf("Employee ID: %d\n Name: %s\n Date of Birth %s\n Position: %s", e.Id, e.Person.Name, e.Person.DateOfBirth, e.Position)
}
