package superstrip

import (
    "strings"
    "fetch/filter"
    "fetch/tokenizer"
)

type Superstrip struct{}

// Handle ascii, lowercase, and stripping punctuation in one filter
func (s *Superstrip) Process(input tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(input, func(token *tokenizer.Token, output tokenizer.TokenChan) {
        cleaned := strings.Map(func(rune int) int {
            switch {
            case 48 <= rune && rune <= 57: // numbers
                fallthrough
            case 97 <= rune && rune <= 122: // lowercase
                return rune
            case 65 <= rune && rune <= 90: // uppercase
                return rune + 32 // Make lowercase
            }
            return -1
        },token.Backing())
        if cleaned != "" {
            output <- tokenizer.NewToken(cleaned)
        }
    })
}

func (s *Superstrip) Cleanup() {}

func Build() *Superstrip {
    return new(Superstrip)
}
