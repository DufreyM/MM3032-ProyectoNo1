# Investigación: LEX, YACC y su Equivalente en Go

---

## ¿Qué es LEX?

**LEX** es una herramienta que se utiliza en compiladores para construir **analizadores léxicos** (también llamados *tokenizadores*).

Un analizador léxico toma una cadena de entrada (como una expresión lógica) y la divide en componentes más pequeños llamados **tokens**, que pueden ser:

- Variables (`p`, `q`, `r`, etc.)
- Operadores (`~`, `^`, `o`, `=>`, `<=>`)
- Paréntesis `(` y `)`
- Constantes (`0`, `1`)

 **Ejemplo:**

Entrada: `(p=>q)^p`  
Salida del lexer: `(`, `p`, `=>`, `q`, `)`, `^`, `p`

---

## ¿Qué es YACC?

**YACC** (*Yet Another Compiler Compiler*) es una herramienta que permite construir **analizadores sintácticos** basados en reglas gramaticales.

Toma los tokens generados por LEX y determina si forman una **estructura válida** según una **gramática definida**.

Además, YACC permite construir un **árbol de análisis sintáctico** (AST), lo cual es útil para interpretar o transformar el contenido de una expresión.

En otras palabras:  
YACC valida la **forma** de la expresión. Se asegura de que los operadores tengan los operandos correctos, los paréntesis estén bien colocados, y que la estructura de la fórmula cumpla con las reglas del lenguaje.

---

## Equivalente en nuestro lenguaje: Go

En este proyecto decidimos utilizar el lenguaje **Go**, que no tiene una herramienta integrada como LEX o YACC. En su lugar, implementamos **nuestro propio analizador léxico (lexer) y sintáctico (parser)** de forma manual, usando estructuras y funciones personalizadas.

>  Esta implementación fue fuertemente influenciada por un proyecto previo desarrollado en el curso de **Estructura de Datos**, donde tuvimos que construir un **intérprete para el lenguaje LISP en Java**. En ese proyecto aprendimos a implementar desde cero un **tokenizador y un lexer**, clasificando símbolos y estructuras según el tipo de expresión. Esa experiencia fue la base conceptual para replicar un flujo similar en Go, pero orientado al reconocimiento de fórmulas del sistema lógico **L**.

---

### Analizador léxico (Lexer)

Creamos una estructura llamada `Lexer` que recorre cada carácter de la cadena de entrada y lo clasifica en distintos tipos de tokens según el alfabeto permitido:

- Letras `p` a `z` → **Variables proposicionales**
- Números `0` y `1` → **Constantes lógicas** (`falso` y `verdadero`)
- Símbolos como `~`, `^`, `o`, `=>`, `<=>`, `(` y `)` → **Operadores lógicos y signos**

Cada vez que el lexer reconoce un símbolo válido, lo devuelve como un **token** que será utilizado por el parser.

---

### Analizador sintáctico (Parser)

Creamos otra estructura llamada `Parser` que implementa las **reglas gramaticales** descritas por el sistema lógico **L**. Estas reglas se aplican a los tokens recibidos por el lexer:

- `expr()` reconoce expresiones de tipo **implicación (`=>`)** y **doble implicación (`<=>`)**
- `term()` reconoce **conjunciones (`^`)** y **disyunciones (`o`)**
- `factor()` reconoce **negaciones (`~`)**, **expresiones entre paréntesis**, **variables** y **constantes**

Cada una de estas funciones construye nodos de un **árbol sintáctico abstracto (AST)**, simulando el comportamiento de un parser generado automáticamente por YACC, pero implementado manualmente.

El **lexer** recorre carácter por carácter la cadena de entrada y clasifica los elementos según el alfabeto del sistema lógico L (variables, operadores, constantes y signos de puntuación), mientras que el **parser** utiliza un enfoque **descendente recursivo** para validar que la estructura cumpla con la gramática especificada, construyendo un **AST (árbol sintáctico abstracto)** que refleja jerárquicamente la fórmula.

---

##  ¿Por qué lo hicimos así?

Go no incluye herramientas automáticas como **LEX/YACC** que sí existen en otros lenguajes como C o Python (por ejemplo, el paquete `PLY`). Por ello, decidimos implementar nuestra propia versión para comprender a profundidad cómo funcionan estos procesos internamente.

Gracias a esta decisión:

- ✅ **Separar** el análisis léxico del análisis sintáctico fue más claro y controlado.
- ✅ Pudimos diseñar un **AST (árbol de sintaxis abstracta)** para representar las expresiones.
- ✅ Pudimos integrar la generación de un **grafo dirigido** visualizable con herramientas como **Graphviz**.

---

## Conclusión

Aunque no usamos herramientas automáticas como LEX y YACC, logramos **replicar su funcionalidad desde cero en Go**. Implementamos un lexer capaz de reconocer tokens válidos del sistema **L** y un parser que valida expresiones según su estructura gramatical.

Este proyecto no solo reforzó nuestro conocimiento sobre **autómatas y gramáticas**, sino que también nos permitió aplicar y consolidar aprendizajes de cursos anteriores como **Estructura de Datos**, conectando ideas de **interpretación de lenguajes** con los fundamentos teóricos de la **lógica matemática**.
