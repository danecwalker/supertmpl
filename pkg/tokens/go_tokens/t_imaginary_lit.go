package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type imaginaryLitToken struct {
	ImaginaryLit string
}

func (t imaginaryLitToken) String() string {
	return t.ImaginaryLit
}

func ImaginaryLit(imaginaryLit string) tokens.Token {
	return imaginaryLitToken{ImaginaryLit: imaginaryLit}
}
