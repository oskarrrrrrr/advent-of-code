package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OpType byte

const (
	OpTypeAdd OpType = '+'
	OpTypeMul OpType = '*'
)

type Expr interface {
	Eval() uint64
}

type Num struct {
	Value uint64
}

func (n Num) Eval() uint64 {
	return n.Value
}

func (n Num) String() string {
	return fmt.Sprint(n.Value)
}

type Op struct {
	Type        OpType
	Left, Right Expr
}

func (o Op) Eval() uint64 {
	switch o.Type {
	case OpTypeAdd:
		return o.Left.Eval() + o.Right.Eval()
	case OpTypeMul:
		return o.Left.Eval() * o.Right.Eval()
	default:
		panic("Unexpected operation type.")
	}
}

func (o Op) String() string {
	return fmt.Sprint("(", o.Left, " ", string(o.Type), " ", o.Right, ")")
}

func ReadInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// We are parsing very simple expressions,
// so separate token types are not really needed.
type TokenStream struct {
	Tokens []byte
	Pos    int
}

func (ts TokenStream) Current() byte {
	return ts.Tokens[ts.Pos]
}

func (ts *TokenStream) Match(chars ...byte) byte {
	for _, c := range chars {
		if ts.Current() == c {
			ts.Advance()
			return c
		}
	}
	return 0
}

func (ts *TokenStream) Advance() {
	ts.Pos++
}

func (ts *TokenStream) Consume(c byte) {
	if ts.Current() == c {
		ts.Advance()
	} else {
		msg := fmt.Sprintf(
			"Expected '%v' but got '%v'.",
			string(c),
			string(ts.Current()),
		)
		panic(msg)
	}
}

func (ts TokenStream) Done() bool {
	return ts.Pos >= len(ts.Tokens)
}

type Precedence int

const (
	PrecedenceEqual = iota
	PrecedenceAddFirst
)

type ParseContext struct {
	Tokens     *TokenStream
	precedence Precedence
}

func (pc ParseContext) parseGrouping() Expr {
	ts := pc.Tokens
	bracket := ts.Match('(')
	if bracket == 0 {
		panic("Could not parse.")
	}
	expr := pc.parseOp()
	ts.Consume(')')
	return expr
}

func (pc ParseContext) parseNum() Expr {
	ts := pc.Tokens
	curr := ts.Current()
	if '0' <= curr && curr <= '9' {
		n, err := strconv.Atoi(string(ts.Current()))
		if err != nil {
			panic(err)
		}
		ts.Advance()
		return Num{uint64(n)}
	}
	return pc.parseGrouping()
}

func (pc ParseContext) parseAnyOp() Expr {
	ts := pc.Tokens
	left := pc.parseNum()
	for !ts.Done() {
		op := OpType(ts.Match('*', '+'))
		if op == 0 {
			return left
		}
		right := pc.parseNum()
		left = Op{Type: op, Left: left, Right: right}
	}
	return left
}

func (pc ParseContext) parseAdd() Expr {
	ts := pc.Tokens
	left := pc.parseNum()
	for !ts.Done() {
		op := OpType(ts.Match('+'))
		if op == 0 {
			return left
		}
		right := pc.parseNum()
		left = Op{Type: '+', Left: left, Right: right}
	}
	return left
}

func (pc ParseContext) parseMul() Expr {
	ts := pc.Tokens
	left := pc.parseAdd()
	for !ts.Done() {
		op := OpType(ts.Match('*'))
		if op == 0 {
			return left
		}
		right := pc.parseAdd()
		left = Op{Type: '*', Left: left, Right: right}
	}
	return left
}

func (pc ParseContext) parseOp() Expr {
	switch pc.precedence {
	case PrecedenceEqual:
		return pc.parseAnyOp()
	case PrecedenceAddFirst:
		return pc.parseMul()
	default:
		panic("Unexpected precedence.")
	}
}

func ParseExpr(s string, precedence Precedence) Expr {
	s = strings.ReplaceAll(s, " ", "")
	tokenStream := &TokenStream{Tokens: []byte(s)}
	parseContext := ParseContext{
		Tokens:     tokenStream,
		precedence: precedence,
	}
	return parseContext.parseOp()
}

func Solve(input string, precedence Precedence) uint64 {
	var result uint64
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		result += ParseExpr(line, precedence).Eval()
	}
	return result
}

func main() {
	input := ReadInput()
	fmt.Println("[1]", Solve(input, PrecedenceEqual))
	fmt.Println("[2]", Solve(input, PrecedenceAddFirst))
}
