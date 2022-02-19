package main

import (
	"fmt"

	"github.com/xenomote/go_parser/pkg/parser"
)

func main() {
	var (
		one = parser.Token{"1", "1"}
		add = parser.Token{"+", "+"}
		// mul = parser.Token{"*", "*"}
	)

	tokens := []parser.Token{one, add, one}
	stream := make(chan parser.Token)

	go writeAll(tokens, stream)

	var (
		e = parser.Symbol{"E", false}
		v = parser.Symbol{"V", false}

		o = parser.Symbol{"1", true}
		a = parser.Symbol{"+", true}
		m = parser.Symbol{"*", true}
	)

	g := parser.Grammar(
		parser.Production("S", e),
		parser.Production("E", e, m, e),
		parser.Production("E", e, a, e),
		parser.Production("E", v),
		parser.Production("V", o),
	)
	p := parser.New(g)

	p.Parse(stream)

	fmt.Printf("%s\n", p.Locations)
	fmt.Println(p.IsMatched())
}

func writeAll(tokens []parser.Token, stream chan<- parser.Token) {
	for _, token := range tokens {
		stream <- token
	}
	close(stream)
}
