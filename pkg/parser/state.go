package parser

import "fmt"

type state struct {
	Production *production
	Index      int
	Origin     int
}

func (s state) IsFinished() bool {
	return s.Index == len(s.Production.Symbols)
}

func (s state) Predicted() Symbol {
	return s.Production.Symbols[s.Index]
}

func (s state) String() string {
	name := s.Production.Name
	left := s.Production.Symbols[:s.Index]
	right := s.Production.Symbols[s.Index:]
	source := s.Origin

	return fmt.Sprintf("(%s->%s.%s, %d)", name, left, right, source)
}

type set struct {
	States []state
}

func (s set) String() string {
	o := "{\n"
	for _, state := range s.States {
		o += fmt.Sprintf("\t%s\n", state)
	}
	o += "}"
	return o
}

func (s *set) Add(x state) {
	for _, other := range s.States {
		if x == other {
			return
		}
	}
	s.States = append(s.States, x)
}