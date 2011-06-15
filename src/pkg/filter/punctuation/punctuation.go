package punctuation

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Punctuation struct {}

func (a *Punctuation) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.BuildFilter(in, func (t *tokenizer.Token) *tokenizer.Token {
        return tokenizer.NewToken(strings.Map(func(rune int) int {
            switch {
            case 48 <= rune && rune <= 57: // numbers
                fallthrough
            case 65 <= rune && rune <= 90: // uppercase
                fallthrough
            case 97 <= rune && rune <= 122: // lowercase
                return rune
            }
            return -1
        }, t.Backing()))
    })
}

func Build() *Punctuation {
    return new(Punctuation)
}
