package punctuation

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Punctuation struct {}

func (p *Punctuation) Process(input tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(input, func(token *tokenizer.Token, output tokenizer.TokenChan) {
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
        }, token.Backing())
        output <- tokenizer.NewToken(cleaned)
    })
}

func (p *Punctuation) Cleanup() {}

func Build() *Punctuation {
    return new(Punctuation)
}
