package main

import (
	"errors"
	"testing"
)

func TestProcesOpcion(t *testing.T) {
	casos := []struct {
		opcionSimulada    int
		espectativaOpcion int
		espectativaBool   bool
		errSimulado       error
	}{
		{5, 5, },
		{7, 7, },
		{-1, -1, },
		{0, 0, },
		{0, 0, true, errors.New("Error Simulado")},
	}
	for _, caso := range casos {
		realBool, opcionSimulada := procesOpcion(caso.opcionSimulada, caso.errSimulado)
		if realBool != caso.espectativaBool {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.espectativaBool, realBool)
		}
		if opcionSimulada != caso.espectativaOpcion {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.espectativaOpcion, opcionSimulada)
		}
	}
}
func TestProcesOperacion(t *testing.T) {
	casos := []struct {
		opcionSimulada    int
		espectativaOpcion int
		
	}{
		{5, 5},
		{7, 7},
		{-1, -1 },
		{0, 0},

	}
	for _, caso := range casos {
		opcionSimulada, err := procesOperacion(caso.opcionSimulada)
		if opcionSimulada != caso.espectativaOpcion {
			t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", caso.espectativaOpcion, opcionSimulada)
	}
		if err := nil{
			t.Error(err)
		}
	}
}
