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
		mostrarTitulo()
		mostrarMenu()
		opcionSolicitada, err := solEntrOpcion()
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
		} else if opcionRecogida <= 0 || opcionRecogida > 6 {
			esperaContinuar()
			limpiarTerminal()
			continue
		}
		operandoInvalido, _ := procesOperacion(opcionRecogida)

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
	cmd := exec.Command("cmd", "/c", "cls", "clear")
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
	case opcionRecogida >= 1 && opcionRecogida <= 6:
		return false, opcionRecogida
	default:
		mostrarInvalido()
		return false, opcionRecogida
	}
}
func procesOperacion(opcionrecogida int) (operandoInvalido bool, resulOpr float32) {

	switch opcionrecogida {
	case 1:
		sumando1, sumando2, _ := solOperAdicion( /*Modificación para PRUEBAS*/ )
		resulOpr = calcAdicion(sumando1, sumando2)
	case 2:
		minuendo, sustraendo, _ := solOperSustraccion()
		resulOpr = calcSustraccion(minuendo, sustraendo)
	case 3:
		factor1, factor2, _ := solOperMultipicacion()
		resulOpr = calcMultiplicacion(factor1, factor2)
	case 4:
		dividendo, divisor, _ := solOperDivision()
		resulOpr = calcDivision(dividendo, divisor)
	case 5:
		base, exponente, _ := solOperPotencia()
		resulOpr = calcPotencia(base, exponente)
	case 6:
		radicando, indice, _ := solOperRaiz()
		resulOpr = calcRaiz(radicando, indice)
	}
	mostrarResulOpr(opcionrecogida, resulOpr)
	return false, resulOpr
}
