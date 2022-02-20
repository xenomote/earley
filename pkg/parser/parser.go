package parser

import "github.com/xenomote/go_parser/pkg/token"

type parser struct {
	Grammar  grammar
	States   []stateSet
	Position int
}

func New(g grammar) parser {
	start := stateSet{}

	for i := range g.Productions {
		production := g.Productions[i]
		if production.Name == g.Initial {
			start.Add(state{&production, 0, 0})
		}
	}

	return parser{
		Grammar:  g,
		States:   []stateSet{start},
		Position: 0,
	}
}

func (p *parser) Parse(ts <-chan token.Token) {
	for token := range ts {
		p.States = append(p.States, stateSet{})
		p.process(token)
		p.Position++
	}
	p.process(token.EOF)
}

func (p *parser) process(t token.Token) {
	// states set may expand during this loop
	for i := 0; i < len(p.current().States); i++ {
		state := p.current().States[i]

		if state.IsFinished() {
			p.match(state)
		} else if state.Predicted().IsTerminal {
			p.scan(state, t)
		} else {
			p.predict(state)
		}
	}
}

func (p *parser) predict(s state) {
	for i := range p.Grammar.Productions {
		production := &p.Grammar.Productions[i]
		if production.Name == s.Predicted().Name {
			prediction := state{production, 0, p.Position}
			p.current().Add(prediction)
		}
	}
}

func (p *parser) scan(s state, t token.Token) {
	if s.Predicted().Name == t.Symbol {
		scanned := s
		scanned.Index++
		p.States[p.Position+1].Add(scanned)
	}
}

func (p *parser) match(s state) {
	origins := p.States[s.Origin]

	for _, origin := range origins.States {
		if origin.Predicted().Name == s.Production.Name {
			matched := origin
			matched.Index++
			p.current().Add(matched)
		}
	}
}

func (p parser) current() *stateSet {
	return &p.States[p.Position]
}

func (p parser) IsMatched() bool {
	for _, state := range p.current().States {
		production := state.Production
		if production.Name == p.Grammar.Initial && state.Index == len(production.Symbols) {
			return true
		}
	}
	return false
}

func (p parser) Trees() []Tree {
	return nil
}