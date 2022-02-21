package parser

import "github.com/xenomote/go_parser/pkg/token"

const INDENT = "  "

type Tree struct {
	Name  string
	Token *token.Token
	Nodes []Tree
}

func (t Tree) String() string {
	return t.stringAux(0)
}

func (t Tree) stringAux(i int) string {
	o := indent(i) + t.Name
	if len(t.Nodes) > 0 {
		o += ":"
	}
	for _, n := range t.Nodes {
		o += "\n" + n.stringAux(i+1)
	}
	return o
}

func indent(i int) string {
	o := ""
	for n := 0; n < i; n++ {
		o += INDENT
	}
	return o
}
