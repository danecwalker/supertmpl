package tokenizer

func NumberState(t *Tokenizer) (StateFn, error) {
	switch t.Peek() {
	case '.':
		return FloatState, nil
	default:
		return IntegerState, nil
	}
}
