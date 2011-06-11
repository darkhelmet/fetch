package ascii

import (
    "fetch/filter"
    "fetch/tokenizer"
)

type Ascii struct {}

func (a *Ascii) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.BuildFilter(in, func (t * tokenizer.Token) *tokenizer.Token {
        return tokenizer.NewToken(t.Backing() + " ascii")
    })
}

func Build() *Ascii {
    return new(Ascii)
}
