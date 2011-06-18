package ascii

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Ascii struct {}

func (a *Ascii) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(in, func(t *tokenizer.Token) []*tokenizer.Token {
        cleaned := strings.Map(func(rune int) int {
            // FIXME: Return Skip() or something
            if rune > 127 {
                return -1
            }
            return rune
        }, t.Backing())
        return []*tokenizer.Token{tokenizer.NewToken(cleaned)}
    })
}

func (a *Ascii) Cleanup() {}

func Build() *Ascii {
    return new(Ascii)
}
