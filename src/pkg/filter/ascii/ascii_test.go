package ascii

import (
    "gospec"
    . "gospec"
    "fetch/tokenizer"
)

func channel(strs ...string) tokenizer.TokenChan {
    filter := Build()
    return filter.Process(tokenizer.NewTokenChan(strs))
}

func AsciiSpec(c gospec.Context) {
    c.Specify("Should remove unicode characters", func() {
        c.Expect((<-channel("Hello 世界")).Backing(), Equals, "Hello ")
    })
}
