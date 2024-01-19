package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type identifierToken struct {
	Identifier string
}

func (t identifierToken) String() string {
	return t.Identifier
}

func Identifier(identifier string) tokens.Token {
	return identifierToken{Identifier: identifier}
}
