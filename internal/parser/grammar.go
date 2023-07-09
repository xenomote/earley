package parser

type grammar struct {
	Initial string
	Rules   []rule
}

func Grammar(p rule, ps ...rule) grammar {
	return grammar{
		Initial: p.Name,
		Rules:   append([]rule{p}, ps...),
	}
}

type rule struct {
	Name       string
	Production production
}

func Rule(name string, symbols ...Symbol) rule {
	return rule{Name: name, Production: symbols}
}

type production []Symbol

type Symbol interface {
	Name() string
	IsTerminal() bool
}

func (p production) String() string {
	o := ""
	for _, s := range p {
		o += s.Name()
	}
	return o
}
