package ascii

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Ascii struct {}

func (a *Ascii) Process(input tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(input, func(token *tokenizer.Token, output tokenizer.TokenChan) {
        cleaned := strings.Map(func(rune int) int {
            if rune > 127 {
                return -1
            }
            return rune
        }, token.Backing())
        output <- tokenizer.NewToken(cleaned)
    })
}

func (a *Ascii) Cleanup() {}

func Build() *Ascii {
    return new(Ascii)
}
