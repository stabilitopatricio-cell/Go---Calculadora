package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Funciones de Entrada
func escanear() {

}

func solicitarOperando1() (float32, error) {
	fmt.Print("\nIntroduce el primer operando: ")
	scanner.Scan()
	entradaNum1 := scanner.Text()
	operando1, err := strconv.ParseFloat(entradaNum1, 32)
	return float32(operando1), err
}
func solicitarOperando2() (float32, error) {
	fmt.Print("\nIntroduce el segundo operando: ")
	scanner.Scan()
	entradaNum2 := scanner.Text()
	operando2, err := strconv.ParseFloat(entradaNum2, 32)
	return float32(operando2), err
}
func solicitarOpcion() (int, error) {
	fmt.Print("Introduce la opcion: ")
	scanner.Scan()
	opcionIngresada := scanner.Text()
	opcionSolicitada, err := strconv.Atoi(opcionIngresada)
	return int(opcionSolicitada), err
}

// Funciones de salida

func mostrarTitulo() {
	fmt.Println(titulo)
}
func mostrarMenu() {
	fmt.Println(menu)
}
func mostrarResultado() {
	fmt.Println("\nEl resutlado es: ", resultado)
}
func mostrarSalida() {
	fmt.Println(mensajeSalida)
}
func mostrarInvalido() {
	fmt.Println(mensajeInvalido)
}

// Información de salida
var titulo string = `
                       >>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<
                          Bienvenido a la calculadora básica --> Go <--
                                       ( v 1.2 by Stapaher )
                       >>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<
`
var menu string = `
                            ==================== O ====================

                            Seleccione la operación que desea realizar:

                               1. Adición             2. Sustracción
                               3. Multiplicación      4. División

                                              0. Salir

                            ==================== O ====================
`
var mensajeSalida string = "\nEspero haber sido de ayuda, !Hasta luego¡\n"
var mensajeInvalido string = "\nOpción no válida - El loro está descansando x_x"

// Información Entrada

var scanner = bufio.NewScanner(os.Stdin)
