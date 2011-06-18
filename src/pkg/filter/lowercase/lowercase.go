package lowercase

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Lowercase struct {}

func (l *Lowercase) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(in, func(t *tokenizer.Token) []*tokenizer.Token {
        return []*tokenizer.Token{tokenizer.NewToken(strings.ToLower(t.Backing()))}
    })
}

func (l *Lowercase) Cleanup() {}

func Build() *Lowercase {
    return new(Lowercase)
}
