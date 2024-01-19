package tokenizer

import gotokens "github.com/danecwalker/supertmpl/pkg/tokens/go_tokens"

func IntegerState(t *Tokenizer) (StateFn, error) {
	for {
		c := t.Peek()
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			t.buf += string(c)
			t.Discard()
		case '.':
			return FloatState, nil
		default:
			t.Emit(gotokens.IntLit(t.buf))
			return SourceState, nil
		}
	}
}
