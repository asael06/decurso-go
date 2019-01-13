package main

import "fmt"

func main() {
	isValid := true
  edad := 14

  if edad < 14 {
    fmt.Println("La edad es menor que 14")
    if edad > 18 {
      fmt.Println("La edad es menor que 18")
    }
  }
	if isValid {
		fmt.Println("Esto está en el bloque de verdadero")
		fmt.Println(isValid)
		if 5 < 10 {
			fmt.Println("5 es menor a 10")
		} else {

		}
fmt.Println("5 no es menor que 10")
	} else 1{
		fmt.Println("Esto es el bloque falso")
		fmt.Println(isValid)
	}
	fmt.Println("Fin del programa")
}
