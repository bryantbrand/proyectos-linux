package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"taller0/matematicas" // Ajusta según tu módulo
)

func main() {
	fmt.Println("¡Hola, Mundo Concurrente!")

	mensaje := "Este es el taller 0"
	imprimirMensaje(mensaje)

	fmt.Println("Contando hasta 3:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Número %d\n", i)
	}

	// --- LECTURA DE ENTRADA DEL USUARIO ---
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Por favor, ingresa un número para calcular su doble: ")
	entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)

	numero, err := strconv.Atoi(entrada)
	if err != nil {
		fmt.Println("Error: Debes ingresar un número válido.", err)
		return
	}

	// --- CÁLCULO DEL DOBLE ---
	resultado := calcularDoble(numero)
	fmt.Printf("El doble de %d es %d\n", numero, resultado)

	// --- CÁLCULO DEL TRIPLE (usando el paquete matematicas) ---
	resultadoTriple := matematicas.CalcularTriple(numero)
	fmt.Printf("El triple de %d es %d\n", numero, resultadoTriple)
}

// Función que imprime un mensaje
func imprimirMensaje(texto string) {
	fmt.Println("Mensaje de la función:", texto)
}

// Función que calcula el doble de un número
func calcularDoble(num int) int {
	return num * 2
}
