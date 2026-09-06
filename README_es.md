
# Calculadora Go — Documentación de evolución
## Versión 0.3 — Separación de responsabilidades
## Identificación de la versión

    Proyecto: Calculadora básica en Go
    Versión: v0.3
    Objetivo principal: separación inicial de responsabilidades
    Estado: funcional y operativo

## Punto de partida

La versión 0.2 había alcanzado un estado funcional estable. El programa permitía:

- Seleccionar una operación.
- Introducir dos operandos.
- Realizar las operaciones básicas.
- Mostrar el resultado.
- Controlar entradas numéricas no válidas.
- Finalizar mediante la opción 0.

La v0.3 no surge para solucionar un error funcional, sino para abordar una cuestión diferente:

> Organizar el código de acuerdo con las responsabilidades que desempeña cada parte del programa.

El objetivo fue comenzar a separar **Flujo principal** | **Interacción con el usuario** | **Operaciones matemáticas**.

## Problema arquitectónico identificado

En la versión anterior, toda la lógica se encontraba concentrada en *main*.

Esto hacía que una única función asumiera simultáneamente responsabilidades de:

- Control del ciclo de ejecución.
- Interacción con el usuario.
- Captura de entradas.
- Conversión de datos.
- Ejecución de operaciones.
- Presentación de resultados.

Aunque el programa funcionaba correctamente, esta concentración dificultaba distinguir las diferentes responsabilidades.

La v0.3 utiliza esta situación como oportunidad para introducir una primera separación estructural.

## Separación de archivos

El código se reorganizó en tres componentes que componen el *paquete main*:

- main.go
- interfazCli.go
- operaciones.go

### main.go

#### Responsabilidad principal:

> Coordinar el flujo de ejecución del programa.

##### Contenido:

- Decisiones sobre el flujo.
- Tratamiento de los errores recibidos.
- Decisiones de coordinación de ejecución de operaciones.

### interfazCli.go

#### Responsabilidad principal:

> Gestionar la interacción entre el usuario y el programa.

###### Contenido:

- Solicitudes de entrada.
- Escaneo de valores introducidos.
- Conversión de entradas.
- Comunicación de valores y errores.
- Funciones de presentación.
- Información textual de la interfaz.

### operaciones.go

#### Responsabilidad principal:

> Realizar las operaciones matemáticas.

###### Contenido:

- Funciones que ejecutan las operaciones.

Estas reciben los operandos y devuelven el resultado correspondiente.

## Evolución del tratamiento de errores

Durante la refactorización apareció una cuestión importante:

    ¿Quién debe decidir qué hacer cuando una entrada produce un error?

Se estableció una separación entre detección y decisión.

### Interfaz

La interfaz realiza la conversión que puede producir el error.

Por ejemplo:

    operando, err := strconv.ParseFloat(entrada, 32)
    return float32(operando), err

La función comunica:

    - Valor obtenido
    - Error producido, si existe.

No decide qué debe hacer el programa con ese error.

### Main

main recibe el resultado y decide cómo continuar:

    operando1, err = solicitarOperando1()

    if err != nil {
        mostrarInvalido()
        continue
    }

La decisión pertenece al flujo de ejecución:

    Mostrar información al usuario -> Descartar la iteración -> Volver a comenzar el ciclo.

### Principio establecido

- Interfaz recibe, detecta, comunica errores de entrada y devuelve información al Usuario
- Main decide que hacer con los datos.

Esta distinción permitió evitar que la interfaz asumiera responsabilidades propias del flujo principal.

## Reducción de estado global

Se eliminó:

> var opcionSolicitada int:

Porque main ya obtiene directamente la opción mediante:

    opcionSolicitada, err := solicitarOpcion()

También se eliminó:

> var opcionIngresada string

    Al comprobar que su función podía resolverse dentro del ámbito local de solicitarOpcion( ).

> Esto representa una mejora respecto a la versión anterior:

    Un dato debe permanecer en el ámbito más reducido posible cuando no necesita ser compartido.

## Responsabilidades consolidadas

### Interfaz

> Solicita -> captura -> convierte -> comunica

### Main

> Recibe -> interpreta -> decide -> coordina

### Operaciones

> Recibe operandos -> calcula -> devuelve resultado

##  Metodología aplicada

La v0.3 permitió evolucionar la metodología utilizada durante el desarrollo.

Inicialmente, el proceso se centraba principalmente en:

    Problema -> Análisis -> Modificación -> Comprobación

Durante esta versión se incorporó una segunda dimensión:

    Funcionamiento -> Análisis de responsabilidades -> Identificación de solapamientos -> Decisión arquitectónica -> Modificación -> Comprobación

Esto supone un cambio importante en el tipo de problemas abordados.

En la v0.2 se trabajó principalmente sobre comportamiento funcional.

En la v0.3 se comenzó a trabajar sobre organización interna del programa.

## Autoevaluación del proceso

Una de las principales conclusiones obtenidas durante esta versión fue que un programa puede funcionar correctamente y, aun así, presentar oportunidades de mejora estructural.

También se comprobó que una refactorización puede generar errores que no existían en el programa original.

Durante la migración aparecieron problemas relacionados con:

- Ámbitos de variables.
- Valores de retorno.
- Tipos incompatibles.
- Nombres duplicados.
- Comunicación entre funciones.
- Replanteamiento de responsabilidades.

Estos problemas no representaban fallos de la lógica original, sino consecuencias de reorganizar su estructura.

El proceso permitió utilizar esos errores como información para comprender mejor el funcionamiento del lenguaje y las responsabilidades de cada componente.

## Elementos deliberadamente pendientes

He identificado cuestiones que no se incorporaron todavía para evitar ampliar innecesariamente el alcance de esta versión.

Estado de variables globales:

    La utilización de:

    var operando1
    var operando2
    var resultado

    como estado compartido queda pendiente de revisión.

### Abstracción de funciones

    solicitarOperando1() y solicitarOperando2()  presentan cierta duplicación.

>No se ha eliminado todavía porque hacerlo introduciría una nueva abstracción que no es necesaria para alcanzar el objetivo principal de esta versión.

    scanner.Scan()
    scanner.Err()

> La aplicación actualmente funciona correctamente bajo el escenario previsto, por lo que esta cuestión no se incorpora como modificación de cierre.

## Resultado de la versión

La v0.3 mantiene el comportamiento funcional de la calculadora y mejora su organización interna.

> El cambio fundamental:

    Límites claros entre las partes del programa.

## Aprendizaje consolidado

El aprendizaje principal de esta versión puede resumirse en una idea:

    La calidad de un programa no depende únicamente de que produzca el resultado correcto, sino también de cómo organiza las responsabilidades necesarias para producirlo.
