package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type semiToken struct{}

func (t semiToken) String() string {
	return ";"
}

func Semi() tokens.Token {
	return semiToken{}
}
