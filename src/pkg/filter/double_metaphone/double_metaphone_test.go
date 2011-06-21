package double_metaphone

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

func DoubleMetaphoneSpec(c gospec.Context) {
    pairs := map[string]string{
        "debugging": "TPKN",
        "laptop":    "LPTP",
        "books":     "PKS",
        "debugger":  "TPKR",
        "winery":    "ANR",
    }
    for original, stemmed := range pairs {
        c.Specify(fmt.Sprintf("Should encode %s to %s", original, stemmed), func() {
            c.Expect((<-channel(original)).Backing(), Equals, stemmed)
        })
    }
}
