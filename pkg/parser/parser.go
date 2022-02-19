package parser

import (
	"fmt"
)

type Symbol struct {
	Name       string
	IsTerminal bool
}

type Symbols []Symbol

func (symbols Symbols) String() string {
	s := ""
	for _, symbol := range symbols {
		s += symbol.Name
	}
	return s
}

type production struct {
	Name    string
	Symbols []Symbol
}

func Production(name string, symbols ...Symbol) production {
	return production{Name: name, Symbols: symbols}
}

type grammar struct {
	Initial     string
	Productions []production
}

func Grammar(initial production, productions ...production) grammar {
	return grammar{
		Initial:     initial.Name,
		Productions: append([]production{initial}, productions...),
	}
}

type Token struct {
	Symbol string
	Source string
}

type State struct {
	Production *production
	Index      int
	Origin     int
}

func (s State) IsFinished() bool {
	return s.Index == len(s.Production.Symbols)
}

func (s State) Predicted() Symbol {
	return s.Production.Symbols[s.Index]
}

func (s State) String() string {
	name := s.Production.Name
	left := Symbols(s.Production.Symbols[:s.Index])
	right := Symbols(s.Production.Symbols[s.Index:])
	source := s.Origin

	return fmt.Sprintf("(%s->%s.%s, %d)", name, left, right, source)
}

type location struct {
	States []State
}

func (l location) String() string {
	s := "{\n"
	for _, state := range l.States {
		s += fmt.Sprintf("\t%s\n", state)
	}
	s += "}"
	return s
}

func (l *location) Add(state State) {
	for _, other := range l.States {
		if state == other {
			return
		}
	}
	l.States = append(l.States, state)
}

type parser struct {
	Grammar   grammar
	Locations []location
	Position  int
}

func New(grammar grammar) parser {
	start := location{}

	for i := range grammar.Productions {
		production := grammar.Productions[i]
		if production.Name == grammar.Initial {
			start.Add(State{&production, 0, 0})
		}
	}

	return parser{
		Grammar:   grammar,
		Locations: []location{start},
		Position:  0,
	}
}

func (p *parser) Parse(tokens <-chan Token) {
	for token := range tokens {
		p.Locations = append(p.Locations, location{})
		p.process(token)
		p.Position++
	}
	p.process(Token{})
}

func (p *parser) process(token Token) {
	for i := 0; i < len(p.Locations[p.Position].States); i++ {
		state := p.Locations[p.Position].States[i]

		if state.IsFinished() {
			p.match(state)
		} else if state.Predicted().IsTerminal {
			p.scan(state, token)
		} else {
			p.predict(state)
		}
	}
}

func (p *parser) predict(state State) {
	for i := range p.Grammar.Productions {
		production := &p.Grammar.Productions[i]
		if production.Name == state.Predicted().Name {
			prediction := State{production, 0, p.Position}
			p.Locations[p.Position].Add(prediction)
		}
	}
}

func (p *parser) scan(state State, token Token) {
	if state.Predicted().Name == token.Symbol {
		scanned := state
		scanned.Index++
		p.Locations[p.Position+1].Add(scanned)
	}
}

func (p *parser) match(state State) {
	origins := p.Locations[state.Origin]

	for _, origin := range origins.States {
		if origin.Predicted().Name == state.Production.Name {
			matched := origin
			matched.Index++
			p.Locations[p.Position].Add(matched)
		}
	}
}

func (p parser) IsMatched() bool {
	for _, state := range p.Locations[p.Position].States {
		production := state.Production
		if production.Name == p.Grammar.Initial && state.Index == len(production.Symbols) {
			return true
		}
	}
	return false
}
