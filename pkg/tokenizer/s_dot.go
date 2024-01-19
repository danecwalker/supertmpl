package tokenizer

import (
	"fmt"
	"unicode"

	gotokens "github.com/danecwalker/supertmpl/pkg/tokens/go_tokens"
)

func DotState(t *Tokenizer) (StateFn, error) {
	t.Discard()
	if unicode.IsDigit(rune(t.Peek())) {
		return NumberState, nil
	}

	switch t.Peek() {
	case '.':
		t.Discard()
		if t.Peek() == '.' {
			t.Emit(gotokens.Operator("..."))
			t.Discard()
			return SourceState, nil
		}
		return nil, fmt.Errorf("unexpected token: %s", string(t.Peek()))
	default:
		t.Emit(gotokens.Operator("."))
		return SourceState, nil
	}
}
