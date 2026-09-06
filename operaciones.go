package main

import (
	"math"
)

// Funciones de operaciones
func calcAdicion(sumando1 float32, sumando2 float32) float32 {
	return sumando1 + sumando2
}
func calcSustraccion(minuendo float32, sustraendo float32) float32 {
	diferencia := minuendo - sustraendo
	return diferencia
}
func calcMultiplicacion(factor1 float32, factor2 float32) float32 {
	producto := factor1 * factor2
	return producto
}
func calcDivision(dividendo float32, divisor float32) float32 {
	cociente := dividendo / divisor
	if divisor == 0 {
		return 0
	} else {
		return cociente
	}
}
func calcPotencia(base float32, exponente float32) float32 {
	potencia := float32(math.Pow(float64(base), float64(exponente)))
	return potencia
}
func calcRaiz(radicando float32, indice float32) float32 {
	raiz := float32(math.Pow(float64(radicando), 1/float64(indice)))
	return float32(raiz)
}
