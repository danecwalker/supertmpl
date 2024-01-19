package tokenizer

import (
	"bufio"
	"io"

	"github.com/danecwalker/supertmpl/pkg/tokens"
)

type Tokenizer struct {
	reader *bufio.Reader
}

func NewTokenizer(reader io.Reader) *Tokenizer {
	return &Tokenizer{
		reader: bufio.NewReader(reader),
	}
}

func (t *Tokenizer) Next() (tokens.Token, error) {
	return nil, nil
}
