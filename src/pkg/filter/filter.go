package filter

import (
    "fetch/tokenizer"
)

type FilterFunc func (token *tokenizer.Token, output tokenizer.TokenChan)
type CleanupFunc func ()

type Filter interface {
    Process(input tokenizer.TokenChan) tokenizer.TokenChan
    Cleanup()
}

func StartFilter(tokens tokenizer.TokenChan, filterToken FilterFunc) tokenizer.TokenChan {
    outc := make(tokenizer.TokenChan, 10)
    go func() {
        for token := range(tokens) {
            filterToken(token, outc)
        }
        close(outc)
    }()
    return outc
}
