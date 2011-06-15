package superstrip

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Superstrip struct {}

// Handle ascii, lowercase, and stripping punctuation in one filter
func (a *Superstrip) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.BuildFilter(in, func (t * tokenizer.Token) (* tokenizer.Token) {
        return tokenizer.NewToken(strings.Map(func(rune int) int {
            switch {
            case 48 <= rune && rune <= 57: // numbers
                fallthrough
            case 65 <= rune && rune <= 90: // uppercase
                return rune + 32 // Make lowercase
            case 97 <= rune && rune <= 122: // lowercase
                return rune
            }
            return -1
        }, t.Backing()))
    })
}

func Build() *Superstrip {
    return new(Superstrip)
}