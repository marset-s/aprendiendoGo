package main

import "fmt"

func main() {
  //array()
  //slice2()
  agregarElementosSlice()
  
}

func array(){
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)
}

func slice(){
  a := make([]int, 5, 10) // len(a) = 5, cap(a) = 10
    fmt.Println(a)          // [0 0 0 0 0]
    fmt.Println(len(a))     // 5
    fmt.Println(cap(a))     // 10

  var s = []bool{true, false}
  fmt.Println(s[0])
}

func slice2() {
  primes := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes[1:4]) // Si no ponemos un valor después de los “:” toma hasta el fin de elementos del slice y viceversa.
}

func agregarElementosSlice(){
  var s []int // Declaramos un slice vacío de enteros
  s = append(s, 2, 3, 4) // Agregamos los elementos 2, 3 y 4 al slice

  fmt.Println(s) // [2 3 4]
}


