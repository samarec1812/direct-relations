// main.go
package main

import (
	"./parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type calcListener struct {
	*parser.BaseRelationListener

	stack []string
	flag bool

}

func (l *calcListener) push(i string) {
	l.stack = append(l.stack, i)
	l.flag = false
}

func (l *calcListener) pop() string {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Pop the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) Contain(s string) bool {
	fl := false
	for i := 0; i < len(l.stack); i++ {
		if l.stack[i] == s {
			fl = true
			break
		}

	}
	return fl
}


// ExitMulDiv is called when exiting the MulDiv production.
func (l *calcListener) ExitCombiner(c *parser.CombinerContext) {
	//right, left := l.pop(), l.pop()

	for i := 0; i < c.GetChildCount(); i++ {

		if (l.Contain(c.GetText()) || strings.Contains(c.GetText(), "(")) { continue }
		l.push(c.GetText())

		//fmt.Println(c.GetRuleContext(), c.GetText())

	}
}


// calc takes a string expression and returns the evaluated result.
func calc(input string) string {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewRelationLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewRelationParser(stream)
	prog := p.Program()
	//	fmt.Println(prog.ToStringTree(lexer.RuleNames,prog.GetParser()))

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	//var c parser.CombinerContext

	antlr.ParseTreeWalkerDefault.Walk(&listener, prog)

	return listener.pop()
}

type node struct {
	node string
	symbol string
}
type edge struct {
	node  string
	symbol string
	label string
}
type graph struct {
	nodes map[node][]edge
}

func NewGraph() *graph {
	return &graph{nodes: make(map[node][]edge)}
}

var numberNode, numberNodeSecond int = 0, 1

func (g *graph) addEdge(from, to, label string) {
	a := 'A' + numberNode
	b := 'A' + numberNodeSecond
	// b := 'А' + numberNode

	g.nodes[node{from, string(a)}] = append(g.nodes[node{node: from, symbol:string(a)}], edge{node: to, symbol:string(b),  label: label})
	if strings.Contains(label, "$end") {
		g.nodes[node{to, string(b)}] = []edge{}
	}
	numberNode++
	numberNodeSecond++
}

func (g *graph) getEdges(node1 node) []edge {
	return g.nodes[node1]
}

func (e *edge) String() string {
	return fmt.Sprintf("%v", e.node)
}

func (g *graph)OutStr() string {
	var out string = ""
	for k := range g.nodes {
		out += "\t"+k.symbol + " " + fmt.Sprintf("[label=\"%s\"];\n", k.node)
	}
	return out
}


func (g *graph) String() string {
	out := `digraph finite_state_machine {
		rankdir=LR;
		size="8,5"
		node [fontname="Arial"; shape = box];
`
	//	mas := g.OutStr(out)
	out += g.OutStr()
	for k := range g.nodes {
		for _, v := range g.getEdges(k) {

			out += fmt.Sprintf("\t%s -> %s\t[ label = \"%s\" ];\n", k.symbol, v.symbol, v.label)
		}
	}
	out += "}"
	return out
}

func main() {
	//fmt.Printf("The answer is: %s\n", calc("(--<#-->).(---) ><#<--"))
	calc("(--<#-->).(---) ><#<--")

	g := NewGraph()
	g.addEdge(`---`, `-->`, "S(a)")
	g.addEdge("-->", "<>", "S($end)")
	fmt.Println(g)

	/*
		g := NewGraph()
		// https://graphviz.gitlab.io/_pages/Gallery/directed/fsm.html

		g.addEdge("LR_0", "LR_2", "SS(B)")
		g.addEdge("LR_0", "LR_1", "SS(S)")
		g.addEdge("LR_1", "LR_3", "S($end)")
		g.addEdge("LR_2", "LR_6", "SS(b)")
		g.addEdge("LR_2", "LR_5", "SS(a)")
		g.addEdge("LR_2", "LR_4", "S(A)")
		g.addEdge("LR_5", "LR_7", "S(b)")
		g.addEdge("LR_5", "LR_5", "S(a)")
		g.addEdge("LR_6", "LR_6", "S(b)")
		g.addEdge("LR_6", "LR_5", "S(a)")
		g.addEdge("LR_7", "LR_8", "S(b)")
		g.addEdge("LR_7", "LR_5", "S(a)")
		g.addEdge("LR_8", "LR_6", "S(b)")
		g.addEdge("LR_8", "LR_5", "S(a)")

		fmt.Println(g)
	*/
}