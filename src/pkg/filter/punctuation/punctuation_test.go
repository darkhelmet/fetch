package punctuation

import (
    "gospec"
    . "gospec"
    "fetch/tokenizer"
)

func channel(strs ...string) tokenizer.TokenChan {
    filter := Build()
    return filter.Process(tokenizer.NewTokenChan(strs))
}

func PunctuationSpec(c gospec.Context) {
    c.Specify("Should remove punctuation", func() {
        c.Expect((<-channel("foobar,./;'[]!@#$^&*()_+")).Backing(), Equals, "foobar")
    })
}
