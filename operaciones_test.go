package main

import (
	"testing"
)

func TestCalcAdicion(t *testing.T) {
	casos := []struct {
		num1     float32
		num2     float32
		pruebaok float32
	}{
		{2, 3, 5},
		{1.8, 1.2, 3},
		{10, -100, -90},
	}
	for _, caso := range casos {
		resultadoPrueba := calcAdicion(caso.num1, caso.num2)
		if resultadoPrueba != caso.pruebaok {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.pruebaok, resultadoPrueba)
		}

	}
}
func TestCalcSustraccion(t *testing.T) {
	casos := []struct {
		num1     float32
		num2     float32
		pruebaok float32
	}{
		{2, 3, -1},
		{1.8, 1, 0.79999995},
		{10, 2, 8},
	}
	for _, caso := range casos {
		resultadoPrueba := calcSustraccion(caso.num1, caso.num2)
		if resultadoPrueba != caso.pruebaok {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.pruebaok, resultadoPrueba)
		}

	}
}
func TestCalcMultiplicacion(t *testing.T) {
	casos := []struct {
		num1     float32
		num2     float32
		pruebaok float32
	}{
		{2, 3, 6},
		{1.8, 2, 3.6},
		{10, 5, 50},
	}
	for _, caso := range casos {
		resultadoPrueba := calcMultiplicacion(caso.num1, caso.num2)
		if resultadoPrueba != caso.pruebaok {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.pruebaok, resultadoPrueba)
		}

	}
}
func TestCalcDivision(t *testing.T) {
	casos := []struct {
		num1     float32
		num2     float32
		pruebaok float32
	}{
		{2, 2, 1},
		{1.8, 2, 0.9},
		{10, 0, 0},
	}
	for _, caso := range casos {
		resultadoPrueba := calcDivision(caso.num1, caso.num2)
		if resultadoPrueba != caso.pruebaok {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.pruebaok, resultadoPrueba)
		}

	}
}
func TestCalcPotencia(t *testing.T) {
	casos := []struct {
		num1     float32
		num2     float32
		pruebaok float32
	}{
		{2, 3, 8},
		{1.8, 1.2, 2.024543},
		{10, -100, 0},
	}
	for _, caso := range casos {
		resultadoPrueba := calcPotencia(caso.num1, caso.num2)
		if resultadoPrueba != caso.pruebaok {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.pruebaok, resultadoPrueba)
		}

	}
}
func TestCalcRaiz(t *testing.T) {
	casos := []struct {
		num1     float32
		num2     float32
		pruebaok float32
	}{
		{52, 3, 3.732511},
		{20, 2, 4.472136},
		{10, -5, 0.63095737},
	}
	for _, caso := range casos {
		resultadoPrueba := calcRaiz(caso.num1, caso.num2)
		if resultadoPrueba != caso.pruebaok {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.pruebaok, resultadoPrueba)
		}

	}
}
