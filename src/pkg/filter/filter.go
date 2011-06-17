package filter

import (
    "fetch/tokenizer"
)

type FilterFunc func (t *tokenizer.Token) *tokenizer.Token
type CleanupFunc func ()

type Filter interface {
    Process(in tokenizer.TokenChan) tokenizer.TokenChan
}

func BuildFilter(in tokenizer.TokenChan, f FilterFunc, cleanup CleanupFunc) tokenizer.TokenChan {
    out := make(tokenizer.TokenChan, 10)
    go func() {
        for token := range(in) {
            nt := f(token)
            if nt != nil {
                out <- nt
            }
        }
        close(out)
        cleanup()
    }()
    return out
}
