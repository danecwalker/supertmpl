package tokenizer

import (
	"fmt"

	gotokens "github.com/danecwalker/supertmpl/pkg/tokens/go_tokens"
)

func StringState(t *Tokenizer) (StateFn, error) {
	for {
		switch t.Peek() {
		case '"':
			t.Discard()
			t.Emit(gotokens.StringLit(t.buf))
			return SourceState, nil
		case '\\':
			t.Discard()
			switch t.Peek() {
			case 'n':
				t.Discard()
				t.buf += "\n"
			case 't':

				t.Discard()
				t.buf += "\t"
			case 'r':
				t.Discard()
				t.buf += "\r"
			case '"':
				t.Discard()
				t.buf += "\""
			case '\\':
				t.Discard()
				t.buf += "\\"
			default:
				return nil, fmt.Errorf("invalid escape sequence")
			}
		case 0:
			return nil, fmt.Errorf("unexpected EOF")
		default:
			t.buf += string(t.Peek())
			t.Discard()
		}
	}
}
