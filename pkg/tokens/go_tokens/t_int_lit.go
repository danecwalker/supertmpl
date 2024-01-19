package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type intLitToken struct {
	IntLit string
}

func (t intLitToken) String() string {
	return t.IntLit
}

func IntLit(intLit string) tokens.Token {
	return intLitToken{IntLit: intLit}
}
