package wisp

import (
	"image/color"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/platinvm/wisp/parser"
)

type visitor struct {
	parser.BaseWispParserVisitor
}

func newVisitor() *visitor {
	return &visitor{}
}

func (v *visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *visitor) VisitProgram(ctx *parser.ProgramContext) any {
	return v.Visit(ctx.Expression())
}

func (v *visitor) VisitExpression(ctx *parser.ExpressionContext) any {
	if ctx.Literal() != nil {
		return v.Visit(ctx.Literal())
	} else if ctx.Array() != nil {
		return v.Visit(ctx.Array())
	} else if ctx.Object() != nil {
		return v.Visit(ctx.Object())
	} else if ctx.Set() != nil {
		return v.Visit(ctx.Set())
	}
	return nil
}

func (v *visitor) VisitArray(ctx *parser.ArrayContext) any {
	values := make([]any, 0)
	for _, expr := range ctx.AllExpression() {
		values = append(values, v.Visit(expr))
	}
	return values
}

func (v *visitor) VisitObject(ctx *parser.ObjectContext) any {
	values := make(map[string]any)
	for _, pair := range ctx.AllPair() {
		p := v.Visit(pair).(struct {
			key   string
			value any
		})
		values[p.key] = p.value
	}
	return values
}

func (v *visitor) VisitSet(ctx *parser.SetContext) any {
	values := make(map[any]struct{})
	for _, expr := range ctx.AllLiteral() {
		values[v.Visit(expr)] = struct{}{}
	}
	return values
}

func (v *visitor) VisitPair(ctx *parser.PairContext) any {
	var key string
	if ctx.ID() != nil {
		key = ctx.ID().GetText()
	} else {
		key = strings.Trim(ctx.STRING().GetText(), `"`)
	}

	value := v.Visit(ctx.Expression())
	return struct {
		key   string
		value any
	}{key, value}
}

func (v *visitor) VisitLiteral(ctx *parser.LiteralContext) any {
	switch {
	case ctx.BOOLEAN() != nil:
		return ctx.BOOLEAN().GetText() == "true"
	case ctx.INTEGER() != nil:
		n, _ := strconv.ParseInt(ctx.INTEGER().GetText(), 10, 64)
		return n
	case ctx.FLOAT() != nil:
		f, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
		return f
	case ctx.BINARY() != nil:
		n, _ := strconv.ParseInt(ctx.BINARY().GetText()[2:], 2, 64)
		return n
	case ctx.HEXADECIMAL() != nil:
		n, _ := strconv.ParseInt(ctx.HEXADECIMAL().GetText()[2:], 16, 64)
		return n
	case ctx.STRING() != nil || ctx.MULTILINE_STRING() != nil:
		return unquote(ctx.GetText())
	case ctx.IPV4() != nil || ctx.IPV6() != nil:
		return net.ParseIP(ctx.GetText())
	case ctx.MAC() != nil:
		hw, _ := net.ParseMAC(ctx.GetText())
		return hw
	case ctx.COLOR() != nil:
		return parseColor(ctx.GetText())
	case ctx.VERSION() != nil:
		return unquote(ctx.GetText())
	case ctx.DURATION() != nil:
		d, _ := time.ParseDuration(ctx.GetText())
		return d
	case ctx.TIMESTAMP() != nil:
		t, _ := time.Parse(time.RFC3339, ctx.GetText())
		return t
	case ctx.PERCENTAGE() != nil:
		s := strings.TrimSuffix(ctx.GetText(), "%")
		p, _ := strconv.ParseFloat(s, 64)
		return p
	}
	return nil
}

func unquote(s string) string {
	s = strings.TrimPrefix(s, "\"")
	s = strings.TrimSuffix(s, "\"")
	return s
}

func parseColor(s string) color.RGBA {
	s = strings.TrimPrefix(s, "#")
	r, _ := strconv.ParseUint(s[0:2], 16, 8)
	g, _ := strconv.ParseUint(s[2:4], 16, 8)
	b, _ := strconv.ParseUint(s[4:6], 16, 8)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func Parse(s string) any {
	input := antlr.NewInputStream(s)
	lexer := parser.NewWispLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewWispParser(stream)

	v := newVisitor()
	return p.Program().Accept(v)
}
