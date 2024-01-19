package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danecwalker/supertmpl/pkg/tokenizer"
)

func main() {
	f, err := os.Open("./examples/basic/basic.stmpl")
	if err != nil {
		log.Fatal(err)
	}

	tk := tokenizer.NewTokenizer(f)

	tokens, err := tk.Scan()
	if err != nil {
		for _, tok := range tokens {
			fmt.Println(tok)
		}
		log.Fatal(err)
	}

	for _, tok := range tokens {
		fmt.Println(tok)
	}
}
