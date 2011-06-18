package lowercase

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Lowercase struct {}

func (l *Lowercase) Process(input tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(input, func(token *tokenizer.Token, output tokenizer.TokenChan) {
        output <- tokenizer.NewToken(strings.ToLower(token.Backing()))
    })
}

func (l *Lowercase) Cleanup() {}

func Build() *Lowercase {
    return new(Lowercase)
}
