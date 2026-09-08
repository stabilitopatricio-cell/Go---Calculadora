package main

import (
	"os"
	"testing"
)

// Prueba de Solicitud de entrada al Usuario
func TestSolEntrUsr(t *testing.T) {
	// Bloqueo de lectura del sistema
	entrOriginal := os.Stdin
	// Reestitución diferida de lectura del sistema
	defer func() { os.Stdin = entrOriginal }()
	// Creación de archivo temporal para entrada simulada
	entrSimulada, err := os.CreateTemp("", "Entrada Simulada")
	if err != nil {
		t.Fatal(err)
	}
	// Eliminación diferida del archivo temporal
	defer os.Remove(entrSimulada.Name())
	// Escritura de texo en archivo temporal
	_, err = entrSimulada.WriteString("Texto de PRUEBA, 1, 2.3, !ª$%&/()=?¿\n")
	if err != nil {
		t.Fatal(err)
	}
	// Reubicación del cursor al inicio de la entrada
	_, err = entrSimulada.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	os.Stdin = entrSimulada

	realidad := solEntrUsr("Solicitud simpulada de entrada : ")
	espectativa := "Texto de PRUEBA, 1, 2.3, !ª$%&/()=?¿"

	if realidad != espectativa {
		t.Errorf(">>> Espectativa : %v <<< | <<< Realidad : %v >>>", espectativa, realidad)
	}
}

// Prueba de Solicitud de Entrada de Opción
func TestSolEntrOpcion(t *testing.T) {
	// Estructura de caso a comprobar
	casos := []struct {
		entrSimulada  string
		espectativa   int
		errorEsperado bool
	}{
		{"3", 3, false},
		{"1.2", 0, true},
		{"abc", 0, true},
	}
	// Bucle de prueba
	for _, caso := range casos {
		// Creación de archivo temporal para almacenar los casos como entradas.
		entrSimulada, err := os.CreateTemp("", "Entrada Simulada")
		if err != nil {
			t.Fatal(err)
		}
		// Eliminación diferida el archivo
		defer os.Remove(entrSimulada.Name())
		// Escritura de casos en archivo temporal
		_, err = entrSimulada.WriteString(caso.entrSimulada + "\n")
		if err != nil {
			t.Fatal(err)
		}
		// Reubicación del cursor al inicio de la entrada simulada
		_, err = entrSimulada.Seek(0, 0)
		if err != nil {
			t.Fatal(err)
		}
		// Redirección de lectura de entrada del sistema a nuestro archivo temporal
		os.Stdin = entrSimulada
		// Llamada de prueba a la función original
		realidad, err := solEntrOpcion()
		if realidad != caso.espectativa {
			t.Errorf(">>> Espectativa : %v <<< | >>> Realidad : %v <<<", caso.entrSimulada, realidad)
		}
		if (err != nil) != caso.errorEsperado {
			t.Errorf("<<< Error inesperado : %v >>>", err)
		}
	}
}

/* Para ejecutar las pruebas de solicitud de operadores para las operaciones se debe adaptar
la función correspondiente en interfazCli y main.
*/
// Prueba de Solicitud de Operandos de Adición >>>>> PENDIENTE DE RESOLVER <<<<<
/*func TestSolOperAdicion(t *testing.T) {
	casos := []struct {
		entrSimulada1 string
		entrSimulada2 string
		espectativa1  float32
		espectativa2  float32
		errorSimulado bool
	}{
		{"2", "3", 2, 3, false},
		{"2.5", "1.75", 2.5, 1.75, false},
		{"abc", "3", 0, 0, true},
		{"2", "abc", 0, 0, true},
	}
	for _, caso := range casos {
		realidad1, realidad2, err := solOperAdicion(caso.entrSimulada1, caso.entrSimulada2)

		if realidad1 != caso.espectativa1 {
			t.Errorf(">>> Espectativa : %v <<< | >>> Realidad : %v <<<", caso.espectativa1, realidad1)
		}
		if realidad2 != caso.espectativa2 {
			t.Errorf(">>> Espectativa : %v <<< | >>> Realidad : %v <<<", caso.espectativa2, realidad2)
		}
		if (err != nil) != caso.errorSimulado {
			t.Error("Error inesperado : ", err)
		}

	}

}*/
