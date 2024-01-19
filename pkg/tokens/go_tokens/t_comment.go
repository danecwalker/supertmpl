package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type commentToken struct {
	Comment string
}

func (t commentToken) String() string {
	return t.Comment
}

func Comment(comment string) tokens.Token {
	return commentToken{Comment: comment}
}
