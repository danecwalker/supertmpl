package tokenizer

type TokenizerError string

func (e TokenizerError) Error() string {
	return string(e)
}

const (
	ErrInvalidToken TokenizerError = "invalid token"
)
