// Code generated from Relation.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Relation

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 34, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 5, 3, 17, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 23, 10, 3, 13, 
	3, 14, 3, 24, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 4, 3, 4, 3, 
	4, 2, 3, 4, 5, 2, 4, 6, 2, 3, 4, 2, 3, 3, 5, 7, 2, 33, 2, 8, 3, 2, 2, 2, 
	4, 16, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 9, 9, 2, 2, 2, 9, 3, 3, 2, 2, 
	2, 10, 11, 8, 3, 1, 2, 11, 17, 7, 4, 2, 2, 12, 13, 7, 8, 2, 2, 13, 14, 
	5, 4, 3, 2, 14, 15, 7, 9, 2, 2, 15, 17, 3, 2, 2, 2, 16, 10, 3, 2, 2, 2, 
	16, 12, 3, 2, 2, 2, 17, 28, 3, 2, 2, 2, 18, 22, 12, 4, 2, 2, 19, 20, 5, 
	2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 23, 3, 2, 2, 2, 22, 19, 3, 2, 2, 2, 23, 
	24, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 
	2, 26, 18, 3, 2, 2, 2, 27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 
	3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 32, 5, 4, 3, 2, 
	32, 7, 3, 2, 2, 2, 5, 16, 24, 28,
}
var literalNames = []string{
	"", "", "", "'#'", "'+'", "'.'", "'('", "')'",
}
var symbolicNames = []string{
	"", "SPACE", "ID_COMBINER", "OPERATOR_PARALLEL", "OPERATOR_SUM", "OPERATOR_POINT", 
	"L_PAR", "R_PAR",
}

var ruleNames = []string{
	"operator", "combiner", "program",
}
type RelationParser struct {
	*antlr.BaseParser
}

// NewRelationParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *RelationParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewRelationParser(input antlr.TokenStream) *RelationParser {
	this := new(RelationParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Relation.g4"

	return this
}


// RelationParser tokens.
const (
	RelationParserEOF = antlr.TokenEOF
	RelationParserSPACE = 1
	RelationParserID_COMBINER = 2
	RelationParserOPERATOR_PARALLEL = 3
	RelationParserOPERATOR_SUM = 4
	RelationParserOPERATOR_POINT = 5
	RelationParserL_PAR = 6
	RelationParserR_PAR = 7
)

// RelationParser rules.
const (
	RelationParserRULE_operator = 0
	RelationParserRULE_combiner = 1
	RelationParserRULE_program = 2
)

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RelationParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RelationParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) OPERATOR_PARALLEL() antlr.TerminalNode {
	return s.GetToken(RelationParserOPERATOR_PARALLEL, 0)
}

func (s *OperatorContext) OPERATOR_SUM() antlr.TerminalNode {
	return s.GetToken(RelationParserOPERATOR_SUM, 0)
}

func (s *OperatorContext) OPERATOR_POINT() antlr.TerminalNode {
	return s.GetToken(RelationParserOPERATOR_POINT, 0)
}

func (s *OperatorContext) SPACE() antlr.TerminalNode {
	return s.GetToken(RelationParserSPACE, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RelationListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RelationListener); ok {
		listenerT.ExitOperator(s)
	}
}




func (p *RelationParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RelationParserRULE_operator)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RelationParserSPACE) | (1 << RelationParserOPERATOR_PARALLEL) | (1 << RelationParserOPERATOR_SUM) | (1 << RelationParserOPERATOR_POINT))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ICombinerContext is an interface to support dynamic dispatch.
type ICombinerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCombinerContext differentiates from other interfaces.
	IsCombinerContext()
}

type CombinerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCombinerContext() *CombinerContext {
	var p = new(CombinerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RelationParserRULE_combiner
	return p
}

func (*CombinerContext) IsCombinerContext() {}

func NewCombinerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CombinerContext {
	var p = new(CombinerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RelationParserRULE_combiner

	return p
}

func (s *CombinerContext) GetParser() antlr.Parser { return s.parser }

func (s *CombinerContext) ID_COMBINER() antlr.TerminalNode {
	return s.GetToken(RelationParserID_COMBINER, 0)
}

func (s *CombinerContext) L_PAR() antlr.TerminalNode {
	return s.GetToken(RelationParserL_PAR, 0)
}

func (s *CombinerContext) AllCombiner() []ICombinerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICombinerContext)(nil)).Elem())
	var tst = make([]ICombinerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICombinerContext)
		}
	}

	return tst
}

func (s *CombinerContext) Combiner(i int) ICombinerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICombinerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICombinerContext)
}

func (s *CombinerContext) R_PAR() antlr.TerminalNode {
	return s.GetToken(RelationParserR_PAR, 0)
}

func (s *CombinerContext) AllOperator() []IOperatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperatorContext)(nil)).Elem())
	var tst = make([]IOperatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperatorContext)
		}
	}

	return tst
}

func (s *CombinerContext) Operator(i int) IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *CombinerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CombinerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CombinerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RelationListener); ok {
		listenerT.EnterCombiner(s)
	}
}

func (s *CombinerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RelationListener); ok {
		listenerT.ExitCombiner(s)
	}
}





func (p *RelationParser) Combiner() (localctx ICombinerContext) {
	return p.combiner(0)
}

func (p *RelationParser) combiner(_p int) (localctx ICombinerContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCombinerContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICombinerContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, RelationParserRULE_combiner, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RelationParserID_COMBINER:
		{
			p.SetState(9)
			p.Match(RelationParserID_COMBINER)
		}


	case RelationParserL_PAR:
		{
			p.SetState(10)
			p.Match(RelationParserL_PAR)
		}
		{
			p.SetState(11)
			p.combiner(0)
		}
		{
			p.SetState(12)
			p.Match(RelationParserR_PAR)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCombinerContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RelationParserRULE_combiner)
			p.SetState(16)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
						{
							p.SetState(17)
							p.Operator()
						}
						{
							p.SetState(18)
							p.combiner(0)
						}




				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(22)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
			}


		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}



	return localctx
}


// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RelationParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RelationParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Combiner() ICombinerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICombinerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICombinerContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RelationListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RelationListener); ok {
		listenerT.ExitProgram(s)
	}
}




func (p *RelationParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RelationParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.combiner(0)
	}



	return localctx
}


func (p *RelationParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
			var t *CombinerContext = nil
			if localctx != nil { t = localctx.(*CombinerContext) }
			return p.Combiner_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RelationParser) Combiner_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

