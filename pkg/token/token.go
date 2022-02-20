package token

var EOF = Token{}

type Token struct {
	Symbol string
	Source string
}

func New(symbol, source string) Token {
	return Token{
		Symbol: symbol,
		Source: source,
	}
}

func Literal(symbol string) Token {
	return Token{
		Symbol: symbol,
		Source: symbol,
	}
}

type stream <-chan Token

func Stream(tokens... Token) stream {
	stream := make(chan Token)
	go writeAll(tokens, stream)
	return stream
}

func writeAll(tokens []Token, stream chan<- Token) {
	for _, token := range tokens {
		stream <- token
	}
	close(stream)
}
