package filter

import (
    "fetch/tokenizer"
)

type FilterFunc func (t *tokenizer.Token) []*tokenizer.Token
type CleanupFunc func ()

type Filter interface {
    Process(in tokenizer.TokenChan) tokenizer.TokenChan
    Cleanup()
}

func StartFilter(tokens tokenizer.TokenChan, f FilterFunc) tokenizer.TokenChan {
    outc := make(tokenizer.TokenChan, 10)
    go func() {
        for token := range(tokens) {
            for _, newToken := range(f(token)) {
                outc <- newToken
            }
        }
        close(outc)
    }()
    return outc
}
