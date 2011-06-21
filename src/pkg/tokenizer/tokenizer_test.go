package tokenizer

import (
    "gospec"
    . "gospec"
)

func TokenizerSpec(c gospec.Context) {
    c.Specify("Should build TokenChan from []string", func() {
        c.Expect((<-NewTokenChan([]string{"foobar"})).Backing(), Equals, "foobar")
    })
}
