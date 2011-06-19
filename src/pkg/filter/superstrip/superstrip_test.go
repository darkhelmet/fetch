package superstrip

import (
    "gospec"
    . "gospec"
    "fetch/tokenizer"
)

func channel(strs... string) tokenizer.TokenChan {
    filter := Build()
    return filter.Process(tokenizer.NewTokenChan(strs))
}

func SuperstripSpec(c gospec.Context) {
    c.Specify("Should trim whitespace", func() {
        for token := range(channel("  foobar  ")) {
            c.Expect(token.Backing(), Equals, "foobar")
        }
    })

    c.Specify("Should convert to lowercase", func() {
        for token := range(channel("LIKEABOSS")) {
            c.Expect(token.Backing(), Equals, "likeaboss")
        }
    })

    c.Specify("Should remove unicode characters", func() {
        for token := range(channel("Hello  世界")) {
            c.Expect(token.Backing(), Equals, "hello")
        }
    })

    c.Specify("Should remove things that aren't letters or numbers", func() {
        for token := range(channel("like,./';'-=a!@#$^*boss")) {
            c.Expect(token.Backing(), Equals, "likeaboss")
        }
    })

    c.Specify("Should keep numbers", func() {
        for token := range(channel("f00b4r")) {
            c.Expect(token.Backing(), Equals, "f00b4r")
        }
    })

    c.Specify("Should not emit empty tokens", func() {
        for token := range(channel("foobar", ",./;'")) {
            c.Expect(token.Backing(), Equals, "foobar")
        }
    })
}
