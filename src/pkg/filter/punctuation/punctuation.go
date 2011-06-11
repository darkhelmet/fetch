package punctuation

import (
    "fetch/filter"
    "fetch/tokenizer"
)

type Punctuation struct {}

func (a *Punctuation) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.BuildFilter(in, func (t * tokenizer.Token) (* tokenizer.Token) {
        return tokenizer.NewToken(t.Backing() + " punctuation")
    })
}

func Build() *Punctuation {
    return new(Punctuation)
}
