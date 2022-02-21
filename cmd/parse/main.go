package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/xenomote/go_parser/pkg/parser"
	"github.com/xenomote/go_parser/pkg/scanner"
	"github.com/xenomote/go_parser/pkg/symbol"
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

	fmt.Printf("%s\n", p.States)
	fmt.Println(p.IsMatched())
}