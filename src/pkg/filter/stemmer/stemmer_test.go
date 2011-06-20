package stemmer

import (
    "fmt"
    "gospec"
    . "gospec"
    "fetch/tokenizer"
)

func channel(strs ...string) tokenizer.TokenChan {
    filter := Build()
    return filter.Process(tokenizer.NewTokenChan(strs))
}

func StemmerSpec(c gospec.Context) {
    pairs := map[string]string{
        "debugging": "debug",
        "laptop":    "laptop",
        "books":     "book",
        "debugger":  "debugg",
        "winery":    "winery",
    }
    for original, stemmed := range pairs {
        c.Specify(fmt.Sprintf("Should stem %s to %s", original, stemmed), func() {
            c.Expect((<-channel(original)).Backing(), Equals, stemmed)
        })
    }
}
