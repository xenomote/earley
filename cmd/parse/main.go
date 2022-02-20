package main

import (
	"fmt"

	"github.com/xenomote/go_parser/pkg/parser"
	"github.com/xenomote/go_parser/pkg/token"
)

func main() {
	var (
		one = token.Literal("1")
		add = token.Literal("+")
		mul = token.Literal("*")
	)

	stream := token.Stream(one, add, one, mul, one)

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

	fmt.Printf("%s\n", p.States)
	fmt.Println(p.IsMatched())

	t := parser.Tree{
		Name: "A",
		Nodes: []parser.Tree{
			{
				Name:  "B",
				Nodes: []parser.Tree{},
			},
			{
				Name: "C",
				Nodes: []parser.Tree{
					{
						Name:  "D",
						Nodes: []parser.Tree{},
					},
				},
			},
		},
	}

	fmt.Println(t)
}
