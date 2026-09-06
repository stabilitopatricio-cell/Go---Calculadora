// Declaración de Paquete
package main

// Importación de paquetes
import (
	// Format
	"fmt"
	// String Conversion
	"strconv"
)

// Función principal - Ejecución
func main() {

	// Impresión del título y menú
	fmt.Println(titulo, menu)
	// Solicitud de opción al Usuario
	fmt.Print("Introduce la opción: ")
	// Captura de opción y guardado en la variable
	fmt.Scanln(&opcionIngresada)
	/* Capa de validación de entrada - Causa: Comportamiento no deseado detectado
	-> Al introducirse un número con decimal, Ej: 1.5, se interpretaba Opción = 1, num1 = 5   */
	opcionSolicitada, err := strconv.Atoi(opcionIngresada)
	if err != nil {
		fmt.Println(`
		Opción no válida - No confundas Churras con Merinas  :S`)
		// Salida de función Madre
		return
	}
	// Condicional de respuesta de salida
	if opcionSolicitada == 0 {
		fmt.Println(`
		Espero haber sido de ayuda, !Hasta luego¡
		`)
		// Condicional de respuesta de opción válida
	} else if opcionSolicitada >= 1 && opcionSolicitada <= 4 {
		fmt.Print(`
		Introduce el primer operando: `)
		fmt.Scanln(&num1)
		fmt.Print(`
		Introduce el segundo operando: `)
		fmt.Scanln(&num2)
		// Condicional de respuesta de opción fuera de rango
	} else {
		fmt.Println(`
		Opción no válida - El loro está descansando x_x
		`)
		// Condicionales de ejecución de operaciones
	}
	if opcionSolicitada == 1 {
		resultadoSumar := num1 + num2
		fmt.Println("El resulado es: ", resultadoSumar)
	} else if opcionSolicitada == 2 {
		resultadoRestar := num1 - num2
		fmt.Println("El resulado es: ", resultadoRestar)
	} else if opcionSolicitada == 3 {
		resultadoMultiplicar := num1 * num2
		fmt.Println("El resulado es: ", resultadoMultiplicar)
	} else if opcionSolicitada == 4 {
		resultadoDividir := num1 / num2
		fmt.Println("El resulado es: ", resultadoDividir)
	}
}

// Título del programa
var titulo string = `
                       >>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<
                          Bienvenido a la calculadora básica --> Go <--
                                       ( v 1.2 by Stapaher )
                       >>>>>>>>>>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<<<<<<<<<
		`

// Menú del Programa
var menu string = `
                            ==================== O ====================

                            Seleccione la operación que desea realizar:

                               1. Adición             2. Sustracción
                               3. Multiplicación      4. División

                                              0. Salir

                            ==================== O ====================
	`

// Opción ingresada por el Usuario
var opcionIngresada string

// Números ingresados para las operaciones
var num1 float32
var num2 float32
