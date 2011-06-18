package filter

import (
    "fetch/tokenizer"
)

type FilterFunc func (t *tokenizer.Token) *tokenizer.Token
type CleanupFunc func ()

type Filter interface {
    Process(in tokenizer.TokenChan) tokenizer.TokenChan
    Cleanup()
}

func StartFilter(inc tokenizer.TokenChan, f FilterFunc) tokenizer.TokenChan {
    outc := make(tokenizer.TokenChan, 10)
    go func() {
        for token := range(inc) {
            nt := f(token)
            if nt != nil {
                outc <- nt
            }
        }
        close(outc)
    }()
    return outc
}
