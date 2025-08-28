package wisp

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/platinvm/wisp/parser"
)

type TreeShapeListener struct {
	*parser.BaseWispParserListener
	Visited []string
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{
		BaseWispParserListener: &parser.BaseWispParserListener{},
		Visited:                []string{},
	}
}

func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	text := ctx.GetText()
	tsl.Visited = append(tsl.Visited, text)
}

// ParseAndVisit parses the input and returns the visited rule texts.
func ParseAndVisit(input string) ([]string, error) {
	is := antlr.NewInputStream(input)
	lexer := parser.NewWispLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewWispParser(stream)
	tree := p.Program()

	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Visited, nil
}
