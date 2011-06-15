package filter

import (
    "fetch/tokenizer"
)

type Filter interface {
    Process(in tokenizer.TokenChan) tokenizer.TokenChan
}

func BuildFilter(in tokenizer.TokenChan, f func (t * tokenizer.Token) (* tokenizer.Token)) tokenizer.TokenChan {
    out := make(tokenizer.TokenChan, 10)
    go func() {
        for token := range(in) {
            out <- f(token)
        }
        close(out)
    }()
    return out
}