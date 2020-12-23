// Code generated from Relation.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 53, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 5, 3, 42, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 
	7, 3, 7, 3, 8, 3, 8, 2, 2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 
	9, 3, 2, 3, 4, 2, 11, 12, 34, 34, 2, 59, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 
	2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7, 43, 
	3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 
	15, 51, 3, 2, 2, 2, 17, 18, 9, 2, 2, 2, 18, 4, 3, 2, 2, 2, 19, 20, 7, 47, 
	2, 2, 20, 21, 7, 47, 2, 2, 21, 42, 7, 47, 2, 2, 22, 23, 7, 47, 2, 2, 23, 
	24, 7, 47, 2, 2, 24, 42, 7, 64, 2, 2, 25, 26, 7, 64, 2, 2, 26, 27, 7, 47, 
	2, 2, 27, 42, 7, 47, 2, 2, 28, 29, 7, 62, 2, 2, 29, 30, 7, 47, 2, 2, 30, 
	42, 7, 47, 2, 2, 31, 32, 7, 47, 2, 2, 32, 33, 7, 47, 2, 2, 33, 42, 7, 62, 
	2, 2, 34, 35, 7, 64, 2, 2, 35, 42, 7, 62, 2, 2, 36, 37, 7, 62, 2, 2, 37, 
	42, 7, 64, 2, 2, 38, 39, 7, 47, 2, 2, 39, 40, 7, 49, 2, 2, 40, 42, 7, 47, 
	2, 2, 41, 19, 3, 2, 2, 2, 41, 22, 3, 2, 2, 2, 41, 25, 3, 2, 2, 2, 41, 28, 
	3, 2, 2, 2, 41, 31, 3, 2, 2, 2, 41, 34, 3, 2, 2, 2, 41, 36, 3, 2, 2, 2, 
	41, 38, 3, 2, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 37, 2, 2, 44, 8, 3, 2, 
	2, 2, 45, 46, 7, 45, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48, 7, 48, 2, 2, 48, 
	12, 3, 2, 2, 2, 49, 50, 7, 42, 2, 2, 50, 14, 3, 2, 2, 2, 51, 52, 7, 43, 
	2, 2, 52, 16, 3, 2, 2, 2, 4, 2, 41, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "'#'", "'+'", "'.'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "SPACE", "ID_COMBINER", "OPERATOR_PARALLEL", "OPERATOR_SUM", "OPERATOR_POINT", 
	"L_PAR", "R_PAR",
}

var lexerRuleNames = []string{
	"SPACE", "ID_COMBINER", "OPERATOR_PARALLEL", "OPERATOR_SUM", "OPERATOR_POINT", 
	"L_PAR", "R_PAR",
}
type RelationLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewRelationLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *RelationLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewRelationLexer(input antlr.CharStream) *RelationLexer {
	l := new(RelationLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Relation.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RelationLexer tokens.
const (
	RelationLexerSPACE = 1
	RelationLexerID_COMBINER = 2
	RelationLexerOPERATOR_PARALLEL = 3
	RelationLexerOPERATOR_SUM = 4
	RelationLexerOPERATOR_POINT = 5
	RelationLexerL_PAR = 6
	RelationLexerR_PAR = 7
)

