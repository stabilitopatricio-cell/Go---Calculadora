package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Funciones de Entrada
func escanear() (float32, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entradaUsr := scanner.Text()
	entrada, err := strconv.ParseFloat(entradaUsr, 32)
	return float32(entrada), err
}
func solicitarOperando1() (float32, error) {
	fmt.Print("\nIntroduce el primer operando: ")
	entrada, err := escanear()
	operando1 := entrada
	return float32(operando1), err
}
func solicitarOperando2() (float32, error) {
	fmt.Print("\nIntroduce el segundo operando: ")
	entrada, err := escanear()
	operando2 := entrada
	return float32(operando2), err
}
func solicitarOpcion() (int, error) {
	fmt.Print("Introduce la opcion: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entradaUsr := scanner.Text()
	opcionSolicitada, err := strconv.Atoi(entradaUsr)
	return int(opcionSolicitada), err
}

// Funciones de salida

func mostrarTitulo() {
	titulo := `
						   >>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<
							  Bienvenido a la calculadora básica --> Go <--
										   ( v 1.2 by Stapaher )
						   >>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<
	`
	fmt.Println(titulo)
}
func mostrarMenu() {
	menu := `
                            ==================== O ====================

                            Seleccione la operación que desea realizar:

                               1. Adición             2. Sustracción
                               3. Multiplicación      4. División

                                              0. Salir

                            ==================== O ====================
	`
	fmt.Println(menu)
}
func mostrarResultado(resultado float32) {
	fmt.Println("\nEl resutlado es: ", resultado)
}
func mostrarSalida() {
	var mensajeSalida string = "\nEspero haber sido de ayuda, !Hasta luego¡\n"
	fmt.Println(mensajeSalida)
}
func mostrarInvalido() {
	var mensajeInvalido string = "\nOpción no válida - El loro está descansando x_x"
	fmt.Println(mensajeInvalido)
}
