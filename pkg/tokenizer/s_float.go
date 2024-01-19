package tokenizer

import (
	"fmt"
	"strings"

	gotokens "github.com/danecwalker/supertmpl/pkg/tokens/go_tokens"
)

func FloatState(t *Tokenizer) (StateFn, error) {
	for {
		c := t.Peek()
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			t.buf += string(c)
			t.Discard()
		case '.':
			if strings.Contains(t.buf, ".") {
				return nil, fmt.Errorf("invalid float")
			}
			t.buf += string(c)
			t.Discard()
		default:
			t.Emit(gotokens.FloatLit(t.buf))
			return SourceState, nil
		}
	}
}
