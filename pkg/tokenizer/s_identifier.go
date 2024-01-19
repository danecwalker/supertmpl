package tokenizer

import (
	"unicode"

	gotokens "github.com/danecwalker/supertmpl/pkg/tokens/go_tokens"
)

func IdentifierState(t *Tokenizer) (StateFn, error) {
	for {
		c := t.Peek()
		if c == 0 {
			break
		}
		if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '_' {
			t.buf += string(c)
			t.Discard()
		} else {
			break
		}
	}

	if _, ok := gotokens.IsKeyword(t.buf); ok {
		t.Emit(gotokens.Keyword(t.buf))
	} else {
		t.Emit(gotokens.Identifier(t.buf))
	}

	return SourceState, nil
}
