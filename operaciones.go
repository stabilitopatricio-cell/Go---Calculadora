package main

// Funciones de operaciones
func sumar(operando1, operando2 float32) float32 {
	return operando1 + operando2
}
func restar(operando1, operando2 float32) float32 {
	return operando1 - operando2
}
func multiplicar(operando1, operando2 float32) float32 {
	return operando1 * operando2
}
func dividir(operando1, operando2 float32) float32 {
	if operando2 <= 0 {
		return 0
	} else {
		return operando1 / operando2
	}
}
