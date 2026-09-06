# Calculadora CLI — v0.4

## Objetivo

La versión v0.4 continúa la evolución estructural de la Calculadora, centrando el desarrollo en la separación de responsabilidades, el alcance de las variables y la mejora progresiva de la organización interna del programa.

No se incorporan nuevas operaciones. El objetivo principal es mejorar la estructura del código y desarrollar criterios de diseño aplicables a futuras versiones.

## Cambios principales
### Separación del procesamiento de opciones

- Se incorpora *procesOpcion( )* para centralizar el procesamiento de la opción introducida por el usuario.

- Se incorpora *procesOperacion( )* para agrupar el procesamiento relacionado con la operación solicitada. 

Actualmente esta función obtiene los operandos y ejecuta la operación correspondiente.

> procesOpcion( ): 

- Recibe la opción solicitada y el posible error asociado.
- Determina si la opción es inválida.
- Comunica la entrada y su invalidez a bucleEjecucion( ).
- Comunica al Usuario (a través de Interfaz) cuando se introduce una opción inválida.


> procesOperacion( ):

- Recibe la opción validada.
- Solicita y valida los operandos y posibles errores asociados.
- Coordina las operaciones solicitadas.
- Comunica al Usuario (a través de Interfaz) el resultado de las operaciones y cuando se introduce un operando inválido.
- Comunica la validéz o invalidéz de las entradas a bucleEjecucion( ).

> bucleEjecucion( ):

- Recibe el resultado de estos procesamientos y decide la acción de control de flujo correspondiente.

### Separación del procesamiento de operaciones

Esta separación constituye una primera aproximación a una organización por responsabilidades y podría evolucionar en versiones posteriores si la complejidad del programa lo requiere.

### Alcance de las variables

Se revisa el ámbito de las variables utilizadas para representar textos de la interfaz.

Los datos que pertenecen exclusivamente a una función pasan progresivamente a declararse como variables locales, evitando mantener como estado global información que no necesita ser compartida.

> Se aplica el criterio:

    Una información debe tener el menor alcance necesario para cumplir su responsabilidad.

### Evaluación de estructuras 

> Cada modificación debe evaluarse según:

- Utilidad real.
- Legibilidad.
- Responsabilidad.
- Mantenibilidad.
- Escalado.

### Conceptos consolidados

- Ámbito de las variables.
- Ciclo de vida de los datos.
- Variables locales frente a variables globales.
- Retorno de múltiples valores.
- Propagación de resultados entre funciones.
- Separación de responsabilidades.
- Control del flujo mediante continue y break.
- Procesamiento y validación de entradas.
- Formateo de valores mediante fmt.
- Especificadores de formato como %.2f.
- Evaluación crítica de sugerencias del compilador y del editor.
- Estado funcional

La calculadora incorpora las operaciones implementadas en versiones anteriores:

- Suma
- Resta
- Multiplicación
- División

>También incorpora:

- Menú interactivo.
- Validación de la opción introducida.
- Salida mediante la opción correspondiente.
- Solicitud de operandos.
- Gestión de entradas inválidas.
- Presentación del resultado.
- Limpieza de la terminal entre iteraciones.

### Criterio de evolución

La v0.4 establece una etapa intermedia entre una implementación funcional y una organización progresivamente más profesional.

Las modificaciones estructurales no se realizan únicamente para reducir líneas de código, sino para conseguir una distribución más clara de responsabilidades y facilitar la evolución futura del programa.

Las decisiones de esta versión quedan cerradas como base para la siguiente etapa de desarrollo.

## Evolución formativa

La versión v0.4 me supone un avance importante no tanto por incorporar nuevas funcionalidades, sino por el cambio en la forma de analizar y estructurar el código.

Durante esta etapa siento que  empiezo a pasar de una visión centrada principalmente en hacer que el programa funcione a una visión más orientada a comprender por qué una determinada estructura es más adecuada que otra.

> Observo una evolución en varios aspectos:

- Mayor atención al alcance de las variables y a su ciclo de vida.
- Comprensión de que una variable global no debe mantenerse como tal si solo pertenece a una función concreta.
- Uso más consciente del paso de información mediante parámetros y valores de retorno.
- Capacidad para distinguir entre una mejora funcional y una mejora estructural.
- Comprensión de que separar responsabilidades puede mejorar mantenibilidad y escalabilidad aunque aumente o no reduzca el número de líneas.
- Mayor criterio para decidir cuándo una abstracción aporta valor y cuándo introduce complejidad innecesaria.
- Evaluación crítica de las sugerencias del editor o del lenguaje, evitando asumir que toda recomendación debe aplicarse automáticamente.
- Mayor atención a la semántica de los nombres y a la legibilidad del código.
- Inicio de una visión más arquitectónica del programa, pensando en cómo podrían evolucionar sus responsabilidades en futuras versiones.
- Ampliación de vocabulario técnico y mayor soltura se expresión.
- Inicio de desarrollo estructural de la documentación.

> También se consolida una forma de razonamiento más cercana al desarrollo profesional:

No basta con preguntarse si el código funciona; también es necesario analizar si sus responsabilidades están bien distribuidas, si los datos tienen el alcance adecuado y si la estructura facilita futuras modificaciones.

Esta versión refleja, por tanto, una transición desde una implementación principalmente funcional hacia una forma de programación progresivamente más consciente del diseño, la mantenibilidad e intención del código.

El objetivo formativo continúa siendo avanzar de manera incremental: primero comprender el fundamento, después aplicarlo en Go y, finalmente, desarrollar criterio para justificar cada decisión de implementación.