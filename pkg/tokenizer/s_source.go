package tokenizer

import gotokens "github.com/danecwalker/supertmpl/pkg/tokens/go_tokens"

func SourceState(t *Tokenizer) (StateFn, error) {
	switch t.Peek() {
	case '+':
		return PlusState, nil
	case '-':
		return MinusState, nil
	case '*':
		return StarState, nil
	case '/':
		return SlashState, nil
	case '%':
		return ModState, nil
	case '&':
		return AndState, nil
	case '|':
		return OrState, nil
	case '^':
		return XorState, nil
	case '<':
		return LessThanState, nil
	case '>':
		return GreaterThanState, nil
	case '=':
		return EqualState, nil
	case '!':
		return BangState, nil
	case ':':
		return ColonState, nil
	case '.':
		return DotState, nil
	case ',':
		t.Emit(gotokens.Operator(","))
		t.Discard()
		return SourceState, nil
	case ';':
		t.Emit(gotokens.Operator(";"))
		t.Discard()
		return SourceState, nil
	case '(':
		t.Emit(gotokens.Operator("("))
		t.Discard()
		return SourceState, nil
	case ')':
		t.Emit(gotokens.Operator(")"))
		t.Discard()
		return SourceState, nil
	case '[':
		t.Emit(gotokens.Operator("["))
		t.Discard()
		return SourceState, nil
	case ']':
		t.Emit(gotokens.Operator("]"))
		t.Discard()
		return SourceState, nil
	case '{':
		t.Emit(gotokens.Operator("{"))
		t.Discard()
		return SourceState, nil
	case '}':
		t.Emit(gotokens.Operator("}"))
		t.Discard()
		return SourceState, nil
	case '~':
		t.Emit(gotokens.Operator("~"))
		t.Discard()
		return SourceState, nil
	default:
		return LiteralState, nil
	}
}
