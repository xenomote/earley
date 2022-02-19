package parser

type Symbol struct {
	Name       string
	IsTerminal bool
}

type symbols []Symbol

func (ss symbols) String() string {
	o := ""
	for _, s := range ss {
		o += s.Name
	}
	return o
}