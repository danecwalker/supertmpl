package tokenizer

import (
	"fmt"
	"unicode"
)

func LiteralState(t *Tokenizer) (StateFn, error) {
	c := t.Peek()
	switch c {
	case '"':
		t.Discard()
		return StringState, nil
	case 0:
		return nil, nil
	default:
		if unicode.IsDigit(rune(c)) {
			return NumberState, nil
		} else if unicode.IsLetter(rune(c)) || c == '_' {
			return IdentifierState, nil
		} else if unicode.IsSpace(rune(c)) {
			t.Discard()
			return SourceState, nil
		} else {
			return nil, fmt.Errorf("unexpected character: %q", c)
		}
	}
}
