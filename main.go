package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// Tokens posibles
const (
	VAR    = "VAR"    
	CONST  = "CONST"  
	NOT    = "~"
	AND    = "^"
	OR     = "o"
	IMPL   = "=>"
	EQ     = "<=>"
	LPAREN = "("
	RPAREN = ")"
	EOF    = "EOF"
)

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	input []rune
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: []rune(input), pos: 0}
}

func (l *Lexer) next() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	ch := l.input[l.pos]
	l.pos++
	return ch
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) NextToken() Token {
	for ch := l.next(); ch != 0; ch = l.next() {
		switch ch {
		case ' ', '\t', '\n':
			continue
		case '(':
			return Token{LPAREN, "("}
		case ')':
			return Token{RPAREN, ")"}
		case '~':
			return Token{NOT, "~"}
		case '^':
			return Token{AND, "^"}
		case 'o':
			return Token{OR, "o"}
		case '=':
			if l.peek() == '>' {
				l.next()
				return Token{IMPL, "=>"}
			}
		case '<':
			if l.peek() == '=' {
				l.next()
				if l.peek() == '>' {
					l.next()
					return Token{EQ, "<=>"}
				}
			}
		default:
			if unicode.IsLetter(ch) && ch >= 'p' && ch <= 'z' {
				return Token{VAR, string(ch)}
			}
			if ch == '0' || ch == '1' {
				return Token{CONST, string(ch)}
			}
			log.Fatalf("token inválido: %q", ch)
		}
	}
	return Token{EOF, ""}
}

type Node interface{}

type BinaryExpr struct {
	Op    string
	Left  Node
	Right Node
}

type UnaryExpr struct {
	Op   string
	Expr Node
}

type Atom struct {
	Value string
}

type Parser struct {
	lexer  *Lexer
	tok    Token
}

func NewParser(input string) *Parser {
	lexer := NewLexer(input)
	return &Parser{lexer: lexer, tok: lexer.NextToken()}
}

func (p *Parser) eat(expected string) {
	if p.tok.Type == expected {
		p.tok = p.lexer.NextToken()
	} else {
		log.Fatalf("se esperaba %s pero se encontró %s", expected, p.tok.Type)
	}
}

func (p *Parser) Parse() Node {
	return p.expr()
}

// expr := term { ("<=>" | "=>") term }
func (p *Parser) expr() Node {
	node := p.term()
	for p.tok.Type == IMPL || p.tok.Type == EQ {
		op := p.tok.Type
		p.eat(op)
		right := p.term()
		node = &BinaryExpr{Op: op, Left: node, Right: right}
	}
	return node
}

// term := factor { ("^" | "o") factor }
func (p *Parser) term() Node {
	node := p.factor()
	for p.tok.Type == AND || p.tok.Type == OR {
		op := p.tok.Type
		p.eat(op)
		right := p.factor()
		node = &BinaryExpr{Op: op, Left: node, Right: right}
	}
	return node
}

// factor := "~" factor | "(" expr ")" | VAR | CONST
func (p *Parser) factor() Node {
	switch p.tok.Type {
	case NOT:
		p.eat(NOT)
		return &UnaryExpr{Op: NOT, Expr: p.factor()}
	case LPAREN:
		p.eat(LPAREN)
		node := p.expr()
		p.eat(RPAREN)
		return node
	case VAR, CONST:
		val := p.tok.Value
		p.eat(p.tok.Type)
		return &Atom{Value: val}
	default:
		log.Fatalf("token inesperado: %v", p.tok)
		return nil
	}
}

var nodeCount int

func generateDOT(n Node) string {
	var b strings.Builder
	b.WriteString("digraph G {\n")
	nodeCount = 0
	visit(n, &b)
	b.WriteString("}")
	return b.String()
}

func visit(n Node, b *strings.Builder) int {
	id := nodeCount
	nodeCount++
	switch t := n.(type) {
	case *Atom:
		fmt.Fprintf(b, "n%d [label=\"%s\"];\n", id, t.Value)
	case *UnaryExpr:
		fmt.Fprintf(b, "n%d [label=\"%s\"];\n", id, t.Op)
		childID := visit(t.Expr, b)
		fmt.Fprintf(b, "n%d -> n%d;\n", id, childID)
	case *BinaryExpr:
		fmt.Fprintf(b, "n%d [label=\"%s\"];\n", id, t.Op)
		leftID := visit(t.Left, b)
		rightID := visit(t.Right, b)
		fmt.Fprintf(b, "n%d -> n%d;\n", id, leftID)
		fmt.Fprintf(b, "n%d -> n%d;\n", id, rightID)
	}
	return id
}

func main() {
	input := "((p=>q)^p)"
	parser := NewParser(input)
	ast := parser.Parse()

	dot := generateDOT(ast)
	os.WriteFile("output.dot", []byte(dot), 0644)
	fmt.Println("Grafo guardado como output.dot (usa Graphviz para visualizar)")
}
