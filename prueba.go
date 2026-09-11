package main

import "fmt"

func main() {
	contador := 1

	// Se ejecuta mientras 'contador' sea menor o igual a 5
	for contador <= 5 {
		fmt.Println("Número:", contador)
		contador++ // Incrementamos para evitar un ciclo infinito
	}
}
