package main

import "fmt"

func main() {
	var numero1, numero2 int
	fmt.Println("Introduce el primer numero: ")
	fmt.Scanln(&numero1)
	fmt.Println("Introduce el segundo numero")
	fmt.Scanln(&numero2)

	fmt.Println("Suma: ", numero1+numero2)
	fmt.Println("Resta:", numero1-numero2)
	fmt.Println("Multiplicación", numero1*numero2)
	fmt.Println("Division", numero1/numero2)
}
