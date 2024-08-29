package main

import "github.com/gin-gonic/gin"


/*
Resumen del flujo

Inicializar el módulo: go mod init nombre-del-modulo

Instalar Gin: go get github.com/gin-gonic/gin

Limpiar dependencias: go mod tidy
*/

func main ()  {

	router := gin.Default()

	//1) Captura la solicitud GET “/hola-mundo”
	router.GET("/hola-mundo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "¡Hola Mundo!",
		})
	})

	//2) Captura la solicitud GET "/ping"
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")

	//Ejecutar con: go run main.go
	//Para parar el servidor hacer: Ctrl + C


}

// Resumen
// El código configura un servidor web básico con Gin que escucha en el puerto 8080.
// Define dos rutas:
// Una que responde a solicitudes GET en /hola-mundo con un mensaje en formato JSON.
// Otra que responde a solicitudes GET en /ping con un simple texto "pong".
// El servidor se ejecuta y queda en espera de recibir solicitudes en el puerto especificado.