package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Funciones de Entrada

func solEntrUsr(mensaje string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(mensaje)
	scanner.Scan()
	entrUsr := scanner.Text()
	return entrUsr
}
func solEntrOpcion() (int, error) {
	entrUsr := solEntrUsr("Introduce la opción: ")
	opcionSolicitada, err := strconv.Atoi(entrUsr)
	if err != nil {
		return 0, err
	}
	return int(opcionSolicitada), nil
}
func solOperAdicion() (float32, float32, error) {
	entrUsr1 := solEntrUsr("\nIntroduce el primer sumando: ")
	sumando1, err := strconv.ParseFloat(entrUsr1, 32)
	if err != nil {
		return 0, 0, err
	}
	entrUsr2 := solEntrUsr("\nIntroduce el segundo sumando: ")
	sumando2, err := strconv.ParseFloat(entrUsr2, 32)
	if err != nil {
		return 0, 0, err
	}
	return float32(sumando1), float32(sumando2), nil
}
func solOperSustraccion() (float32, float32, error) {
	entrUsr1 := solEntrUsr("\nIntroduce el minuendo: ")
	minuendo, err := strconv.ParseFloat(entrUsr1, 32)
	if err != nil {
		return 0, 0, err
	}
	entrUsr2 := solEntrUsr("\nIntroduce el sustraendo: ")
	sustraendo, err := strconv.ParseFloat(entrUsr2, 32)
	if err != nil {
		return 0, 0, err
	}
	return float32(minuendo), float32(sustraendo), nil
}
func solOperMultipicacion() (float32, float32, error) {
	entrUsr1 := solEntrUsr("\nIntroduce el primer factor: ")
	factor1, err := strconv.ParseFloat(entrUsr1, 32)
	if err != nil {
		return 0, 0, err
	}
	entrUsr2 := solEntrUsr("\nIntroduce el segundo factor: ")
	factor2, err := strconv.ParseFloat(entrUsr2, 32)
	if err != nil {
		return 0, 0, err
	}
	return float32(factor1), float32(factor2), nil
}
func solOperDivision() (float32, float32, error) {
	entrUsr1 := solEntrUsr("\nIntroduce el dividendo: ")
	dividendo, err := strconv.ParseFloat(entrUsr1, 32)
	if err != nil {
		return 0, 0, err
	}
	entrUsr2 := solEntrUsr("\nIntroduce el divisor: ")
	divisor, err := strconv.ParseFloat(entrUsr2, 32)
	if err != nil {
		return 0, 0, err
	}
	return float32(dividendo), float32(divisor), nil
}
func solOperPotencia() (float32, float32, error) {
	entrUsr1 := solEntrUsr("\nIntroduce el base: ")
	base, err := strconv.ParseFloat(entrUsr1, 32)
	if err != nil {
		return 0, 0, err
	}
	entrUsr2 := solEntrUsr("\nIntroduce el exponente: ")
	exponente, err := strconv.ParseFloat(entrUsr2, 32)
	if err != nil {
		return 0, 0, err
	}
	return float32(base), float32(exponente), nil
}
func solOperRaiz() (float32, float32, error) {
	entrUsr1 := solEntrUsr("\nIntroduce el radicando: ")
	radicando, err := strconv.ParseFloat(entrUsr1, 32)
	if err != nil {
		return 0, 0, err
	}
	entrUsr2 := solEntrUsr("\nIntroduce el índice: ")
	indice, err := strconv.ParseFloat(entrUsr2, 32)
	if err != nil {
		return 0, 0, err
	}
	return float32(radicando), float32(indice), nil
}

// Funciones de salida

func mostrarTitulo() {
	titulo := `
>>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<
    Bienvenido a la calculadora básica --> Go <--
                ( v 1.0 by Stapaher )
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
        5. Potenciación        6. Radicación
                    0. Salir

    ==================== O ====================
	`
	fmt.Println(menu)
}
func mostrarResulOpr(opcionSolicitada int, resulOpr float32) {
	switch opcionSolicitada {
	case 1:
		fmt.Println("\nEl resultado es: ", resulOpr)
	case 2:
		fmt.Println("\nLa diferencia es: ", resulOpr)
	case 3:
		fmt.Println("\nEl producto es: ", resulOpr)
	case 4:
		fmt.Println("\nEl cociente es: ", resulOpr)
	case 5:
		fmt.Println("\nLa potencia es: ", resulOpr)
	case 6:
		fmt.Println("\nLa raíz es: ", resulOpr)
	}
}
func mostrarSalida() {
	var mensajeSalida string = "\nEspero haber sido de ayuda, !Hasta luego¡\n"
	fmt.Println(mensajeSalida)
}
func mostrarInvalido() {
	var mensajeInvalido string = "\nOpción no válida - El loro está descansando x_x"
	fmt.Println(mensajeInvalido)
}
