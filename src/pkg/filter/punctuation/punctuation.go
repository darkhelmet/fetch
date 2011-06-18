package punctuation

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Punctuation struct {}

func (p *Punctuation) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(in, func(t *tokenizer.Token) []*tokenizer.Token {
        cleaned := strings.Map(func(rune int) int {
            switch {
            case 48 <= rune && rune <= 57: // numbers
                fallthrough
            case 65 <= rune && rune <= 90: // uppercase
                fallthrough
            case 97 <= rune && rune <= 122: // lowercase
                return rune
            }
            return -1
        }, t.Backing())
        return []*tokenizer.Token{tokenizer.NewToken(cleaned)}
    })
}

func (p *Punctuation) Cleanup() {}

func Build() *Punctuation {
    return new(Punctuation)
}
