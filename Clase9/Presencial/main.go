package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go mod init
//go get -u github.com/gin-gonic/gin
//go tidy

//gin tiene 2 middleware que vamos a estar utilizando: Logger y Recovery

//ejercicio: crear una persona en formato json y devolver por consola
//primero creamos nuestra estrcutura

type Persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Direccion string
	Telefono  string
	Activo    bool
}

func main (){

router := gin.Default() //exporta los loger y recovery por defecto

var p Persona

//sacamos los tildes para que no falle el unmarchal👀
jsonData := `{"Nombre":"Juan", "Apellido":"Perez", "Edad": 25, "Direccion":"Calle falsa 123", "Telefono":"1155447788", "Activo": true}`

//queremos parsear el array de bytes y parsearlo a mi estrcutura persona

if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
	log.Fatal(err)
}

fmt.Println(p) //devuelve en formato go: estructura


//persona es mi ruta definidad para url /persona
router.GET("/persona", func(ctx *gin.Context){
	panic("test")
	ctx.JSON(http.StatusOK, gin.H{
		"persona": p, //devuelve en formato json, al cliente
	})
})

router.GET("/persona1", func(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"persona": p, 
	})
})

router.Run(":8080")

}


