# Calculadora básica — Go ( by Stapaher )

## 1. Descripción

Ésta en mi *Calculadora básica* desarrollada en **Go** como primer proyecto práctico de programación.

El proyecto tiene como objetivo materializar los conceptos fundamentales de programación mediante una aplicación de consola funcional, aplicando progresivamente principios de organización, separación de responsabilidades, flujo de datos y control de ejecución.

**Versión actual:** `v1.0`
**Estado:** Finalizada

---

## 2. Funcionalidades

> La aplicación permite realizar seis operaciones matemáticas básicas:

1. Adición
2. Sustracción
3. Multiplicación
4. División
5. Potenciación
6. Radicación

Y finalizar la ejecución desde el menú principal.

---

## 3. Características

* Funciona mediante interfaz de línea de comandos.
* Solicita y valida la opción seleccionada por el usuario.
* Solicita los operandos correspondientes a cada operación.
* Convierte las entradas de texto a los tipos numéricos necesarios.
* Controla las entradas no válidas.
* Mantiene el flujo principal mediante un bucle de ejecución.
* Separa la entrada, el procesamiento, el cálculo y la presentación.
* Utiliza funciones específicas para cada responsabilidad.
* Utiliza `math.Pow()` para las operaciones de potenciación y radicación.

---

## 4. Estructura funcional

La aplicación está organizada conceptualmente en cuatro áreas principales:

### Entrada

> Responsable de obtener información introducida por el usuario.

```text
solEntrUsr()
solEntrOpcion()
solOperAdicion()
solOperSustraccion()
solOperMultipicacion()
solOperDivision()
solOperPotencia()
solOperRaiz()
```

La entrada general se obtiene como `string`. La conversión al tipo correspondiente se realiza posteriormente según el significado de los datos.

### Procesamiento

> Responsable de interpretar las opciones y dirigir la ejecución.

```text
procesOpcion()
procesOperacion()
```

### Cálculo

> Responsable de realizar las operaciones matemáticas.

```text
calcAdicion()
calcSustraccion()
calcMultiplicacion()
calcDivision()
calcPotencia()
calcRaiz()
```

### Salida

> Responsable de presentar información al usuario.

```text
mostrarTitulo()
mostrarMenu()
mostrarResulOpr()
mostrarSalida()
mostrarInvalido()
```

---

## 5. Flujo general de ejecución

El flujo principal de la aplicación es:

```text
main()
  ↓
bucleEjecucion()
  ↓
mostrarTitulo()
  ↓
mostrarMenu()
  ↓
solEntrOpcion()
  ↓
procesOpcion()
  ↓
procesOperacion()
  ↓
solOper...
  ↓
calc...
  ↓
mostrarResulOpr()
  ↓
esperaContinuar()
  ↓
limpiarTerminal()
  ↓
siguiente iteración
```

La función `bucleEjecucion()` mantiene el ciclo de vida de la aplicación y determina cuándo continuar o finalizar.

---

## 6. Operaciones matemáticas

### Adición

Recibe dos sumandos y devuelve la suma.

```text
sumando1 + sumando2 → suma
```

### Sustracción

Recibe minuendo y sustraendo y devuelve la diferencia.

```text
minuendo - sustraendo → diferencia
```

### Multiplicación

Recibe dos factores y devuelve el producto.

```text
factor1 × factor2 → producto
```

### División

Recibe dividendo y divisor y devuelve el cociente.

```text
dividendo ÷ divisor → cociente
```

La división por cero se controla antes de devolver el resultado.

### Potenciación

Recibe base y exponente.

```text
base^exponente → potencia
```

Se implementa mediante `math.Pow()`.

### Radicación

Recibe radicando e índice.

```text
ⁿ√radicando → raíz
```

Se implementa mediante la equivalencia matemática:

```text
ⁿ√x = x^(1/n)
```

utilizando `math.Pow()`.

---

## 7. Control del flujo

El programa utiliza un bucle `for` para mantener activa la calculadora.

El control del flujo bucle es la única responsabilidad de *bucleEjecucion( )*:

* `continue`: reinicia la iteración actual.
* `break`: finaliza el bucle.


Las funciones auxiliares proporcionan información al bucle para que este pueda determinar el flujo que debe seguir.

---

## 8. Gestión de errores

Las entradas del usuario se reciben inicialmente como texto.

Posteriormente se convierten mediante:

```go
strconv.Atoi()
```

para las opciones del menú y:

```go
strconv.ParseFloat()
```

para los valores numéricos.

Los errores de conversión se propagan mediante valores `error` y son tratados antes de continuar con la operación correspondiente.

---

## 9. Tipos numéricos

Las operaciones matemáticas de la calculadora utilizan `float32`.

Las funciones de `math` utilizadas para potenciación y radicación trabajan con `float64`, por lo que se realizan las conversiones necesarias:

```text
float32
   ↓
float64
   ↓
math.Pow()
   ↓
float64
   ↓
float32
```

---

## 10. Organización del código

El proyecto se organiza actualmente un único paquete, con 3 archivos según su responsabilidad:

```text
main.go
interfazCli.go
operaciones.go
```

### `main.go`

Contiene el punto de entrada y el control principal de ejecución.

### `interfazCli.go`

Contiene las funciones relacionadas con la relación Programa <-> Usuario:

* Entrada de usuario.
* Solicitud de operandos.
* Presentación de información.
* Interfaz de consola.

### `operaciones.go`

Contiene las funciones encargadas de realizar los cálculos matemáticos.

---

## 11. Evolución del proyecto

El desarrollo comenzó con una implementación sencilla orientada principalmente a conseguir que la calculadora funcionara.

Durante las sucesivas versiones se fueron introduciendo mejoras estructurales a medida que aparecían nuevas necesidades.

La evolución llevó progresivamente a:

```text
Código funcional
      ↓
Separación de responsabilidades
      ↓
Flujo de datos explícito
      ↓
Control del ámbito y ciclo de vida
      ↓
Funciones con responsabilidades concretas
      ↓
Estructura preparada para nuevas operaciones
```

La refactorización realizada durante las versiones anteriores no tuvo como objetivo reducir el número de líneas, sino mejorar la organización y la comprensión del código.

---

## 12. Evolución formativa

Este proyecto constituye el primer ejercicio práctico completo dentro del proceso de aprendizaje de programación con **Go**.

La evolución del proyecto ha permitido trabajar progresivamente conceptos como:

* variables y tipos;
* funciones;
* parámetros y valores de retorno;
* ámbito de las variables;
* ciclo de vida de los datos;
* estructuras de control;
* bucles;
* `switch`;
* `continue` y `break`;
* conversión de tipos;
* gestión de errores;
* entrada y salida;
* separación de responsabilidades;
* flujo explícito de información;
* nomenclatura semántica;
* refactorización;
* evaluación crítica de sugerencias del entorno de desarrollo.

El aprendizaje no se ha limitado a conseguir un resultado funcional. También se ha trabajado el criterio necesario para determinar **por qué una determinada estructura resulta adecuada para una necesidad concreta**.

Una decisión válida en una versión anterior puede dejar de ser la más adecuada cuando cambian los requisitos. La refactorización se utiliza, por tanto, como herramienta para adaptar la estructura del programa a las nuevas necesidades.

---

## 13. Alcance de la v1.0

La versión `v1.0` se considera finalizada al disponer de una calculadora básica funcional que integra las seis operaciones previstas y una estructura suficientemente organizada para los objetivos del primer proyecto.

No se incorporan en esta versión funcionalidades adicionales como:

* Expresiones matemáticas combinadas.
* Precedencia entre operadores.
* Operaciones con un número variable de operandos.
* Historial de operaciones.
* Persistencia de datos.
* Interfaz gráfica.
* Nuevas categorías de operaciones.

Estas funcionalidades podrán constituir objetivos de versiones futuras.

---

## 14. Estado

**Versión:** `v1.0`
**Estado:** Finalizada
**Tipo:** Primer proyecto práctico de programación con Go

