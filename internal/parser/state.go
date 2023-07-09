package parser

import (
	"fmt"
)

type state struct {
	Rule   *rule
	Index  int
	Origin int
}

func (s state) IsFinished() bool {
	return s.Index == len(s.Rule.Production)
}

func (s state) Predicted() Symbol {
	return s.Rule.Production[s.Index]
}

func (s state) String() string {
	name := s.Rule.Name
	left := s.Rule.Production[:s.Index]
	right := s.Rule.Production[s.Index:]
	source := s.Origin

	return fmt.Sprintf("(%s->%s.%s, %d)", name, left, right, source)
}

type stateSet struct {
	States []state
}

func (s *stateSet) Add(x state) {
	for _, other := range s.States {
		if x == other {
			return
		}
	}
	s.States = append(s.States, x)
}

func (s stateSet) String() string {
	o := "{\n"
	for _, state := range s.States {
		o += fmt.Sprintf("\t%s\n", state)
	}
	o += "}"
	return o
}
