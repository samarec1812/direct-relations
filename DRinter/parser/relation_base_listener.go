// Code generated from Relation.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Relation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRelationListener is a complete listener for a parse tree produced by RelationParser.
type BaseRelationListener struct{}

var _ RelationListener = &BaseRelationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRelationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRelationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRelationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRelationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseRelationListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseRelationListener) ExitOperator(ctx *OperatorContext) {}

// EnterCombiner is called when production combiner is entered.
func (s *BaseRelationListener) EnterCombiner(ctx *CombinerContext) {}

// ExitCombiner is called when production combiner is exited.
func (s *BaseRelationListener) ExitCombiner(ctx *CombinerContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRelationListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRelationListener) ExitProgram(ctx *ProgramContext) {}
