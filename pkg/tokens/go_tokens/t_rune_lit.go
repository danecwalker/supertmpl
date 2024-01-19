package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type runeLitToken struct {
	RuneLit string
}

func (t runeLitToken) String() string {
	return t.RuneLit
}

func RuneLit(runeLit string) tokens.Token {
	return runeLitToken{RuneLit: runeLit}
}
