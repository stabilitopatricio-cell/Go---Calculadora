package main

func main() {
	bucleEjecucion()
}

func bucleEjecucion() {
	for {
		mostrarTitulo()
		mostrarMenu()
		opcionSolicitada, err := solicitarOpcion()
		if err != nil {
			mostrarInvalido()
			continue
		}
		if opcionSolicitada == 0 {
			mostrarSalida()
			break
		} else if opcionSolicitada >= 1 && opcionSolicitada <= 4 {
			operando1, err = solicitarOperando1()
			if err != nil {
				mostrarInvalido()
				continue
			}
			operando2, err = solicitarOperando2()
			if err != nil {
				mostrarInvalido()
				continue
			}
		} else {
			mostrarInvalido()
			continue
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
		mostrarResultado()
	}
}
