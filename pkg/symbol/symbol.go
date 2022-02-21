package symbol

type symbol struct {
	name       string
	isTerminal bool
}

func (s symbol) Name() string {
	return s.name
}

func (s symbol) IsTerminal() bool {
	return s.isTerminal
}

func Terminal(s string) symbol {
	return symbol{ name: s, isTerminal: true }
}

func NonTerminal(s string) symbol {
	return symbol{ name: s, isTerminal: false }
}
