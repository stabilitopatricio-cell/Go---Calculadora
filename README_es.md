# Introducción 

Durante el desarrollo de mi formación autodidacta en programación, me hé abrumado ante la cantidad de información, opciones, variantes, ámbitos de desarrollo, etc; debido a mi costrumbre a un sistema educativo modularmente estructurado, continúo buscando el 'punto de anclaje' para mi evolución.

Ahora se trata del leguaje Go, en una busqueda casi desesperada de simplicidad para la consolidación de conocimiento.

## Sobre el proyecto

Esta es la primera etapa del proyecto formativo en lenguaje GO.

La inteción es un programa básico que irá evolicionado en complejidad
tanto como la sintaxis estudiada de forma paralela en el momento del desarrollo del mismo.

## Desarrollo

Separaré el proyecto el *versiones funcionales*, en el que cada versión representa un cambio perceptible de complejidad a nivel conceptual.

Primera etapa: 

- Comprender estructuras sintácticas básicas características del lenguaje 
    * Cadenas de caracteres
    * Comentarios
    * Operadores
    * Condicionales
    * Importadiones
    * Funciones
    * Etc

### Comportamientos no deseados detectados

- Al introducirse un número con decimal, cuando se solicita una *opción int* de operación.
    * Ej: 1.5, *fmt.Scannln* interpreta -> *opción = 1, num1 = 5* -> Cuando debería ser *Opción no válida*
        > Gestionado con *strconv.Atoi()*. 

- Al solicitar *num1 y num2*, si se introduce un *float*, se interpreta como dos entras consecutivas.
    * Ej: Usuario introcuce 1.35 como *num1*, para la *opción 3 (Multiplicación)* -> *fmt.Scannln* interpreta, *num1 = 1, num2 = 35*.
        > Solucionado designando ambas variable como *float32*.

- Ahora al introducir un *String* durante la *solicitud de opción*, se ejecuta *if opción == 0*.
    * Si durante la *solicitud de opción* se introduce un *string*, Ej: Hola, continua la ejecución de *if == 0*.
        > Solucionado con la implementación de *return*.

