package main

import (
	"bufio"
	"os"
	"os/exec"
)

func main() {

	bucleEjecucion()
}

func bucleEjecucion() {

	for {
		var operando1, operando2 float32
		mostrarTitulo()
		mostrarMenu()
		opcionSolicitada, err := solicitarOpcion()
		opcionInvalida, opcionRecogida := procesOpcion(opcionSolicitada, err)
		if opcionInvalida == true {
			esperaContinuar()
			limpiarTerminal()
			continue
		}
		if opcionRecogida == 0 {
			esperaContinuar()
			limpiarTerminal()
			break
		} else if opcionRecogida <= 0 || opcionRecogida > 4 {
			esperaContinuar()
			limpiarTerminal()
			continue
		}
		operandoInvalido := procesOperacion(opcionRecogida, operando1, operando2, err)

		if operandoInvalido == true {
			esperaContinuar()
			limpiarTerminal()
			continue
		} else {
			esperaContinuar()
			limpiarTerminal()
		}

	}
}
func esperaContinuar() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
func limpiarTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func procesOpcion(opcionSolicitada int, err error) (bool, int) {
	opcionRecogida, opcionInvalida := opcionSolicitada, err
	switch {
	case opcionInvalida != nil:
		mostrarInvalido()
		return true, opcionRecogida
	case opcionRecogida == 0:
		mostrarSalida()
		return false, opcionRecogida
	case opcionRecogida >= 1 && opcionRecogida <= 4:
		return false, opcionRecogida
	default:
		mostrarInvalido()
		return false, opcionRecogida
	}
}
func procesOperacion(opcionSolicitada int, operando1 float32, operando2 float32, err error) (operandoInvalido bool) {
	var resultado float32
	if err != nil {
		mostrarInvalido()
		return true
	} else {
		operando1, err = solicitarOperando1()
		operando2, err = solicitarOperando2()
	}
	switch opcionSolicitada {
	case 1:
		resultado = sumar(operando1, operando2)
	case 2:
		resultado = restar(operando1, operando2)
	case 3:
		resultado = multiplicar(operando1, operando2)
	case 4:
		resultado = dividir(operando1, operando2)

	}
	mostrarResultado(resultado)
	return false
}
