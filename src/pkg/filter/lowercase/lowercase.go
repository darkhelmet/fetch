package lowercase

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Lowercase struct {}

func (a *Lowercase) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.BuildFilter(in, func (t * tokenizer.Token) (* tokenizer.Token) {
        return tokenizer.NewToken(strings.ToLower(t.Backing()))
    })
}

func Build() *Lowercase {
    return new(Lowercase)
}
