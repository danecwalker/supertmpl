package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type stringLitToken struct {
	StringLit string
}

func (t stringLitToken) String() string {
	return t.StringLit
}

func StringLit(stringLit string) tokens.Token {
	return stringLitToken{StringLit: stringLit}
}
