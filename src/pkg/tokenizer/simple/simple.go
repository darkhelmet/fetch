package simple

import (
    "fetch/tokenizer"
    "strings"
)

type SimpleTokenizer struct{}

func (st *SimpleTokenizer) Tokenize(input string) tokenizer.TokenChan {
    return tokenizer.NewTokenChan(strings.Fields(input))
}

func Build() *SimpleTokenizer {
    return new(SimpleTokenizer)
}
