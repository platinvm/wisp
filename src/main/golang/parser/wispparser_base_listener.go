// Code generated from WispParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // WispParser

import "github.com/antlr4-go/antlr/v4"

// BaseWispParserListener is a complete listener for a parse tree produced by WispParser.
type BaseWispParserListener struct{}

var _ WispParserListener = &BaseWispParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWispParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWispParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWispParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWispParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseWispParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseWispParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWispParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWispParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArray is called when production array is entered.
func (s *BaseWispParserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseWispParserListener) ExitArray(ctx *ArrayContext) {}

// EnterObject is called when production object is entered.
func (s *BaseWispParserListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseWispParserListener) ExitObject(ctx *ObjectContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseWispParserListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseWispParserListener) ExitPair(ctx *PairContext) {}

// EnterSet is called when production set is entered.
func (s *BaseWispParserListener) EnterSet(ctx *SetContext) {}

// ExitSet is called when production set is exited.
func (s *BaseWispParserListener) ExitSet(ctx *SetContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseWispParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseWispParserListener) ExitLiteral(ctx *LiteralContext) {}
