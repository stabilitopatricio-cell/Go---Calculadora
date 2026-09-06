
# Análisis de evolución — Calculadora Go v0.2

## Punto de partida

> La v0.1 tenía un objetivo deliberadamente limitado:

- Conseguir una calculadora básica funcional.

> La aplicación permitía:

    - Seleccionar una operación.
    - Introducir dos operandos.
    - Realizar suma, resta, multiplicación y división.
    - Salir mediante la opción 0.

> No era una implementación ideal. Su propósito era demostrar que el programa podía funcionar de extremo a extremo.

> La v0.2 parte de esa base sin intentar todavía refactorizarla completamente.

## Problema principal detectado

> Durante la utilización de la *v0.1* apareció un comportamiento no deseado al introducir una opción con varios caracteres.

- Por ejemplo:

    La entrada *1.5* no era tratada como una única unidad por *fmt.Scanln* en el contexto utilizado.

    Esto provocaba que *1.5* terminara siendo procesado parcialmente y que el valor restante afectara a la siguiente lectura.

> El problema no era simplemente:

    "No se convertir 1.5."

> El problema fundamental era que:

- El mecanismo utilizado para capturar la entrada no estaba proporcionando el modelo de entrada que necesitábamos.

> Esta distinción fue determinante para encontrar la solución.

## Cambio fundamental
### fmt.Scanln → bufio.Scanner

> Introduje:

    scanner := bufio.NewScanner(os.Stdin)

> y posteriormente:

    scanner.Scan()
    opcionIngresada = scanner.Text()*

> El cambio modifica el modelo de entrada:

- ANTES:

    entrada -> *Scanln* -> interpretación/conversión

- AHORA:

    entrada -> *Scanner* -> *string* -> *strconv* -> tipo
---
### Impacto en el funcionamiento

> Una entrada como:

    1.5

> se captura como una única línea:

    "1.5"

> Ya no depende de que *Scanln* vaya consumiendo parcialmente los elementos de la entrada.

> Esto resuelve el problema funcional que originó la investigación.

## Separación entre captura y conversión

> Este ha sido probablemente el cambio conceptual más importante de la versión.

> La entrada del usuario se captura primero como texto:

    entradaNum1 := scanner.Text()

> y posteriormente se interpreta:

    operando1, err := strconv.ParseFloat(entradaNum1, 32)

> Por tanto:

    CAPTURA:

    Scanner -> string

    INTERPRETACIÓN:

    strconv -> float32

> Esto introduce una separación clara de responsabilidades.

### Impacto

> Ahora el programa puede distinguir entre:

    - ¿Qué ha introducido el usuario?
    - ¿Cómo se interpreta esa entrada?
    - ¿La interpretación ha sido posible?

> Esto será especialmente importante cuando aumente la complejidad del programa.

## Introducción práctica de error

> Ya había implementado previamente:

    numero, err := strconv.Atoi(...)

> En esta versión el concepto se aplica de forma práctica también a los operandos:

    operando1, err := strconv.ParseFloat(entradaNum1, 32)

> y:

    if err != nil {
        ...
        continue
    }

> Esto permite detectar una entrada como:

    abc

> sin provocar que el programa continúe utilizando un valor numérico inválido.

### Impacto

> La aplicación adquiere una primera capa explícita de validación.

> No se trata todavía de un sistema completo de validación, pero sí de una mejora importante respecto a la v0.1.

### Introducción de continue como mecanismo de recuperación

> Ya tenía incorporado *for* como bucle principal de ejecución.

> En esta versión *continue* adquiere una función práctica:

    if err != nil {
        fmt.Println(...)
        continue
    }

> La aplicación puede ahora:

    entrada incorrecta -> mostrar mensaje -> descartar iteración actual -> volver al menú

> Esto permite que un error de entrada no termine el programa.

## Uso de switch

> La ejecución de las operaciones pasó de una cadena de *if / else if* a:

    switch opcionSolicitada {
    case 1:
        ...
    case 2:
        ...
    case 3:
        ...
    case 4:
        ...
    }

### Impacto

> No añade una capacidad nueva al programa, pero mejora la correspondencia entre la estructura del código y su propósito:

    opción 1 → operación 1
    opción 2 → operación 2
    opción 3 → operación 3
    opción 4 → operación 4

> También prepara el código para una futura ampliación del número de operaciones.

> La refactorización completa de esta estructura, no obstante, queda deliberadamente fuera de esta versión.

## Metodología aplicada

> La metodología que he intentado aplicar ha sido algo similar a:
 
    PRUEBA DE FUNCIONAMIENTO
        ↓
    PROBLEMA
        ↓
    OBSERVACIÓN
        ↓
    EXPERIMENTACIÓN
        ↓
    HIPÓTESIS
        ↓
    NUEVA EXPERIMENTACIÓN
        ↓
    REFORMULACIÓN DEL PROBLEMA
        ↓
    DOCUMENTACIÓN OFICIAL
        ↓
    COMPARACIÓN DE ALTERNATIVAS
        ↓
    ELECCIÓN
        ↓
    IMPLEMENTACIÓN
        ↓
    PRUEBA

### Observación

> Durante la prueba de funcionamiento se advierten comportamientos concretos:

    1.5 → comportamiento inesperado

> Al observar comoportamientos no deseados, hé realizado deliberadamente pruebas adicionales para provocar todos los comoportamientos a corregir y abordarlos en conjunto:

    g
    gh
    1.5
    Hola
    dh,sdjo.123
    1.2,4.r,f
    etc

> Esto me confirmó que el problema era más amplio que un caso concreto.

### Reformulación

> La primera interpretación me sugirió:

    "Tengo un problema de conversión."

> Pero el análisis posterior me permitió especificar mejor:

    Tengo un problema relacionado con cómo se captura y delimita la entrada.

> Esto cambia directamente el enfoque.

### Comparación de alternativas

> Hé analizado la documentación oficial de <https://pkg.go.dev/std>. Específicamente los paquetes *fmt, bufio, errors, io, strconv, string, stucts y builtin*.

> Y lo que se acercaba mas a mi situación era:

    fmt.Scanln
    strconv
    bufio.Reader
    bufio.Scanner

> Luego se especificaron sus responsabilidades.

> La conclusión fue:

    Scanner -> capturar
    strconv -> convertir

> en lugar de intentar utilizar una única herramienta para resolver ambas cuestiones.

### Utilización de documentación oficial

> Esto permitió comprobar específicamente el funcionamiento de bufio.Scanner, ScanLines, Text, Scan, Err, Reader, etc.

> La documentación dejó de ser simplemente una fuente para **buscar una función** y pasó a utilizarse para **evaluar** alternativas de diseño.

## Evolución de la formación

> La formación también ha avanzado junto al proyecto.

> En la v0.1 predominaba:

    ¿Cómo hago que funcione?

> En la v0.2 empieza a aparecer:

    ¿Qué está haciendo realmente el programa?

> El proyecto empieza a dejar de ser solamente un ejercicio de sintaxis Go y comienza a funcionar como instrumento para aprender y pulir la  metodología aplicada.


