package ascii

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Ascii struct {}

func (a *Ascii) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.BuildFilter(in, func (t * tokenizer.Token) *tokenizer.Token {
        return tokenizer.NewToken(strings.Map(func(rune int) int {
            if rune > 127 {
                return -1
            }
            return rune
        }, t.Backing()))
    })
}

func Build() *Ascii {
    return new(Ascii)
}
