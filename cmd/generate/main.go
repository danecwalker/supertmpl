package main

import (
	"log"
	"os"

	"github.com/danecwalker/supertmpl/pkg/tokenizer"
)

func main() {
	f, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}

	tk := tokenizer.NewTokenizer(f)
}
