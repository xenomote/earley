package token

var EOF = Token{}

type Token struct {
	symbol string
	source string
}

func (t Token) Symbol() string {
	return t.symbol
}

func New(symbol, source string) Token {
	return Token{
		symbol: symbol,
		source: source,
	}
}

func Literal(s string) Token {
	return Token{
		symbol: s,
		source: s,
	}
}

func Stream(ts... Token) <-chan Token {
	stream := make(chan Token)
	go writeAll(ts, stream)
	return stream
}

func writeAll(ts []Token, s chan<- Token) {
	for _, t := range ts {
		s <- t
	}
	close(s)
}
