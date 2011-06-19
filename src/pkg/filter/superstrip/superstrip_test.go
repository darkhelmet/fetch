package superstrip

import (
    "gospec"
    . "gospec"
    "fetch/tokenizer"
)

func channel(strs ...string) tokenizer.TokenChan {
    filter := Build()
    return filter.Process(tokenizer.NewTokenChan(strs))
}

func SuperstripSpec(c gospec.Context) {
    c.Specify("Should trim whitespace", func() {
        c.Expect((<-channel("  foobar  ")).Backing(), Equals, "foobar")
    })

    c.Specify("Should convert to lowercase", func() {
        c.Expect((<-channel("LIKEABOSS")).Backing(), Equals, "likeaboss")
    })

    c.Specify("Should remove unicode characters", func() {
        c.Expect((<-channel("Hello  世界")).Backing(), Equals, "hello")
    })

    c.Specify("Should remove things that aren't letters or numbers", func() {
        c.Expect((<-channel("like,./';'-=a!@#$^*boss")).Backing(), Equals, "likeaboss")
    })

    c.Specify("Should keep numbers", func() {
        c.Expect((<-channel("f00b4r")).Backing(), Equals, "f00b4r")
    })

    c.Specify("Should not emit empty tokens", func() {
        for token := range channel("foobar", ",./;'") {
            c.Expect(token.Backing(), Equals, "foobar")
        }
    })
}
