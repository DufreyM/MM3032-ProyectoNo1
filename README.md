# Investigación: LEX, YACC y su Equivalente en Go

---

## ¿Qué es LEX?

**LEX** es una herramienta que se utiliza en compiladores para construir **analizadores léxicos** (también llamados *tokenizadores*).

Un analizador léxico toma una cadena de entrada (como una expresión lógica) y la divide en componentes más pequeños llamados **tokens**, que pueden ser:

* Variables (`p`, `q`, `r`, etc.)
* Operadores (`~`, `^`, `o`, `=>`, `<=>`)
* Paréntesis `(` y `)`
* Constantes (`0`, `1`)

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

> Esta implementación fue fuertemente influenciada por un proyecto previo desarrollado en el curso de **Estructura de Datos**, donde tuvimos que construir un **intérprete para el lenguaje LISP en Java**. En ese proyecto aprendimos a implementar desde cero un **tokenizador y un lexer**, clasificando símbolos y estructuras según el tipo de expresión. Esa experiencia fue la base conceptual para replicar un flujo similar en Go, pero orientado al reconocimiento de fórmulas del sistema lógico **L**.

---

### Analizador léxico (Lexer)

Creamos una estructura llamada `Lexer` que recorre cada carácter de la cadena de entrada y lo clasifica en distintos tipos de tokens según el alfabeto permitido:

* Letras `p` a `z` → **Variables proposicionales**
* Números `0` y `1` → **Constantes lógicas** (`falso` y `verdadero`)
* Símbolos como `~`, `^`, `o`, `=>`, `<=>`, `(` y `)` → **Operadores lógicos y signos**

Cada vez que el lexer reconoce un símbolo válido, lo devuelve como un **token** que será utilizado por el parser.

---

### Analizador sintáctico (Parser)

Creamos otra estructura llamada `Parser` que implementa las **reglas gramaticales** descritas por el sistema lógico **L**. Estas reglas se aplican a los tokens recibidos por el lexer:

* `expr()` reconoce expresiones de tipo **implicación (`=>`)** y **doble implicación (`<=>`)**
* `term()` reconoce **conjunciones (`^`)** y **disyunciones (`o`)**
* `factor()` reconoce **negaciones (`~`)**, **expresiones entre paréntesis**, **variables** y **constantes**

Cada una de estas funciones construye nodos de un **árbol sintáctico abstracto (AST)**, simulando el comportamiento de un parser generado automáticamente por YACC, pero implementado manualmente.

El **lexer** recorre carácter por carácter la cadena de entrada y clasifica los elementos según el alfabeto del sistema lógico L (variables, operadores, constantes y signos de puntuación), mientras que el **parser** utiliza un enfoque **descendente recursivo** para validar que la estructura cumpla con la gramática especificada, construyendo un **AST (árbol sintáctico abstracto)** que refleja jerárquicamente la fórmula.

---

## Gramática usada (Sistema L)

Las reglas implementadas siguen la gramática del sistema L:

* Las variables proposicionales y constantes del alfabeto de L son fórmulas bien formadas.
* Si A es una fórmula bien formada, entonces `~A` también lo es.
* Si A y B son fórmulas bien formadas, entonces `A^B`, `AoB`, `A=>B`, `A<=>B` también lo son.
* Si A es una fórmula bien formada, entonces `(A)` también lo es.

Esto se representa en el código con funciones anidadas:

* `factor()` = reconoce `~`, `VAR`, `CONST`, `(` `expr` `)`
* `term()` = reconoce operadores binarios de menor precedencia: `^`, `o`
* `expr()` = reconoce operadores binarios de mayor precedencia: `=>`, `<=>`

---

## Generación del AST (Grafo dirigido)

El AST se genera con la función `generateDOT()` en el archivo `main.go`, que recorre el árbol generado por el parser y genera un archivo `.dot` para visualizarlo con Graphviz.

Ejemplo para la fórmula `(p=>q)^p`:

```
    ^
   / \
 =>   p
/  \
p    q
```

---
<img width="200" height="200" alt="image" src="https://github.com/user-attachments/assets/97003992-6a5f-49e2-99b1-f26bb5ed47a7" />

---

## Visualización

El archivo generado `output.dot` puede abrirse con [Graphviz](https://graphviz.org/) para visualizar el árbol sintáctico. Usa el siguiente comando para obtener la imagen:

```bash
dot -Tpng output.dot -o arbol.png
```

---

## Bibliografía

* Aho, A. V., Lam, M. S., Sethi, R., & Ullman, J. D. (2006). *Compilers: Principles, Techniques, and Tools*. Pearson Education.
* [Documentación oficial de Graphviz](https://graphviz.org/)
* [PLY (Python Lex-Yacc)](https://www.dabeaz.com/ply/)
* [Go Programming Language Specification](https://golang.org/ref/spec)
* Ejemplo y adaptación basada en proyectos previos del curso de Estructura de Datos (2024).
