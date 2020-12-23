// Code generated from Relation.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Relation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RelationListener is a complete listener for a parse tree produced by RelationParser.
type RelationListener interface {
	antlr.ParseTreeListener

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterCombiner is called when entering the combiner production.
	EnterCombiner(c *CombinerContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitCombiner is called when exiting the combiner production.
	ExitCombiner(c *CombinerContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)
}
