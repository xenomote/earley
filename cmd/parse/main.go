package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xenomote/earley/internal/parser"
	"github.com/xenomote/earley/internal/scanner"
	"github.com/xenomote/earley/internal/symbol"
)

func main() {
	var (
		e = symbol.NonTerminal("E")
		v = symbol.NonTerminal("V")

		o = symbol.Terminal("1")
		a = symbol.Terminal("+")
		m = symbol.Terminal("*")
	)

	s := scanner.New(
		scanner.Pattern(o, regexp.MustCompile("0|[1-9][0-9]*")),
	)
	ts := s.Scan(strings.NewReader("1 + 1 * 3"))

	g := parser.Grammar(
		parser.Rule("S", e),
		parser.Rule("E", e, m, e),
		parser.Rule("E", e, a, e),
		parser.Rule("E", v),
		parser.Rule("V", o),
	)
	p := parser.New(g)

	p.Parse(ts)

	fmt.Println(p.States)
	fmt.Println(p.IsMatched())
}
