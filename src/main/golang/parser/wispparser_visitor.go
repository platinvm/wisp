// Code generated from WispParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // WispParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by WispParser.
type WispParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by WispParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by WispParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by WispParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by WispParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by WispParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by WispParser#set.
	VisitSet(ctx *SetContext) interface{}

	// Visit a parse tree produced by WispParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
