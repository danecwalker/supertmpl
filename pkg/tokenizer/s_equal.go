package tokenizer

import gotokens "github.com/danecwalker/supertmpl/pkg/tokens/go_tokens"

func EqualState(t *Tokenizer) (StateFn, error) {
	t.Discard()
	switch t.Peek() {
	case '=':
		t.Emit(gotokens.Operator("=="))
		t.Discard()
		return SourceState, nil
	default:
		t.Emit(gotokens.Operator("="))
		return SourceState, nil
	}
}
