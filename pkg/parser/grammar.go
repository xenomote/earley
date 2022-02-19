package parser

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

type production struct {
	Name    string
	Symbols symbols
}

func Production(name string, symbols ...Symbol) production {
	return production{Name: name, Symbols: symbols}
}