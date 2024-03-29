package scanner

import (
	"io"
	"regexp"

	"github.com/xenomote/earley/internal/token"
)

type scanner struct {
	Location location
}

type location struct {
	row, column int
}

func New(ps ...pattern) scanner {
	return scanner{
		Location: location{row: 0, column: 0},
	}
}

func (s *scanner) Scan(r io.RuneReader) <-chan token.Token {
	panic("not implemented")
}

type pattern struct {
	match  *regexp.Regexp
	symbol symbol
}

type symbol interface {
	Name() string
}

func Pattern(s symbol, r *regexp.Regexp) pattern {
	return pattern{symbol: s, match: r}
}
