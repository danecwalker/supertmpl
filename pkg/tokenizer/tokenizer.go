package tokenizer

import (
	"bufio"
	"io"

	"github.com/danecwalker/supertmpl/pkg/tokens"
)

type Tokenizer struct {
	reader      *bufio.Reader
	state       StateFn
	tokenStream []tokens.Token
	buf         string
}

type StateFn func(*Tokenizer) (StateFn, error)

func NewTokenizer(reader io.Reader) *Tokenizer {
	return &Tokenizer{
		reader:      bufio.NewReader(reader),
		state:       SourceState,
		tokenStream: make([]tokens.Token, 0),
	}
}

func (t *Tokenizer) Scan() (toks []tokens.Token, err error) {
	for t.state != nil {
		t.state, err = t.state(t)
		if err != nil {
			return t.tokenStream, err
		}
	}

	return t.tokenStream, nil
}

func (t *Tokenizer) Peek() byte {
	c, err := t.reader.Peek(1)
	if err != nil {
		return 0
	}

	return c[0]
}

func (t *Tokenizer) Discard() error {
	_, err := t.reader.Discard(1)
	return err
}

func (t *Tokenizer) Emit(tok tokens.Token) {
	t.tokenStream = append(t.tokenStream, tok)
	t.buf = ""
}

func (t *Tokenizer) LastToken() tokens.Token {
	return t.tokenStream[len(t.tokenStream)-1]
}
