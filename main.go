// Declaración de Paquete
package main

// Importación de paquetes
import (
	// Operative System
	"os"
	// Format
	"fmt"
	// Buffered Input-Output
	"bufio"
	// String Conversion
	"strconv"
)

// Función principal - Ejecución
func main() {
	// Búffer de entradas
	scanner := bufio.NewScanner(os.Stdin)
	// Bucle de ejecución
	for {
		// Impresión del título y menú
		fmt.Println(titulo, menu)
		// Solicitud de opción al Usuario
		fmt.Print("Introduce la opcion: ")
		// Captura de opción y guardado en la variable
		scanner.Scan()
		opcionIngresada = scanner.Text()
		/* Capa de validación de entrada - Causa: Comportamiento no deseado detectado
		-> Al introducirse un número con decimal, Ej: 1.5, se interpretaba Opción = 1, num1 = 5   */
		opcionSolicitada, err := strconv.Atoi(opcionIngresada)
		if err != nil {
			fmt.Println(`
									Opción no válida - No confundas Churras con Merinas  :S
									`)
			// Reinicio del bucle de ejecución
			continue
		}
		// Condicional de respuesta de salida
		if opcionSolicitada == 0 {
			fmt.Println(`
			Espero haber sido de ayuda, !Hasta luego¡
			`)
			break
			// Condicional de respuesta de opcion válida
		} else if opcionSolicitada >= 1 && opcionSolicitada <= 4 {
			// Soicitud de operandos al Usuario, captura de entradas y guardado en variable
			fmt.Print(`
				Introduce el primer operando: `)
			scanner.Scan()
			entradaNum1 := scanner.Text()
			operando1, err := strconv.ParseFloat(entradaNum1, 32)
			if err != nil {
				fmt.Println("Alto Compi, te noto desmejorao. - Prueba otra vez ;)")
				continue
			}
			num1 = float32(operando1)

			fmt.Print(`
			Introduce el segundo operando: `)
			scanner.Scan()
			entradaNum2 := scanner.Text()
			operando2, err := strconv.ParseFloat(entradaNum2, 32)
			if err != nil {
				fmt.Println("!Para Tron¡ Que siento consquillas - Prueba otra vez, anda ;)")
				continue
			}
			num2 = float32(operando2)

			// Condicional de respuesta de opcion fuera de rango
		} else {
			fmt.Println(`
			Opción no válida - El loro está descansando x_x
			`)
			continue
		}
		// Condicionales de ejecución de operaciones
		switch opcionSolicitada {
		case 1:
			resultadoSumar := num1 + num2
			fmt.Println("El resulado es: ", resultadoSumar)
		case 2:
			resultadoRestar := num1 - num2
			fmt.Println("El resulado es: ", resultadoRestar)
		case 3:
			resultadoMultiplicar := num1 * num2
			fmt.Println("El resulado es: ", resultadoMultiplicar)
		case 4:
			resultadoDividir := num1 / num2
			fmt.Println("El resulado es: ", resultadoDividir)
		}
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
