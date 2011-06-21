package stopword

import (
    "gospec"
    . "gospec"
    "fetch/tokenizer"
)

func channel(strs ...string) tokenizer.TokenChan {
    filter := Build()
    return filter.Process(tokenizer.NewTokenChan(strs))
}

func keys(words map[string]bool) []string {
    keys := make([]string, len(words))
    i := 0
    for key, _ := range words {
        keys[i] = key
        i++
    }
    return keys
}

func StopwordSpec(c gospec.Context) {
    c.Specify("Should ignore stopwords", func() {
        stopwords := keys(words)
        for token := range channel(stopwords...) {
            c.Expect(token.Backing(), Equals, "should never get here")
        }
    })
}
