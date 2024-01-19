package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type floatLitToken struct {
	FloatLit string
}

func (t floatLitToken) String() string {
	return t.FloatLit
}

func FloatLit(floatLit string) tokens.Token {
	return floatLitToken{FloatLit: floatLit}
}
