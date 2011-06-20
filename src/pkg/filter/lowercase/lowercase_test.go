package lowercase

import (
    "gospec"
    . "gospec"
    "fetch/tokenizer"
)

func channel(strs ...string) tokenizer.TokenChan {
    filter := Build()
    return filter.Process(tokenizer.NewTokenChan(strs))
}

func LowercaseSpec(c gospec.Context) {
    c.Specify("Should convert to lowercase", func() {
        c.Expect((<-channel("FOOBAR")).Backing(), Equals, "foobar")
    })
}
