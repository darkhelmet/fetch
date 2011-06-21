package simple

import (
    "gospec"
    . "gospec"
)

func SimpleSpec(c gospec.Context) {
    c.Specify("Should tokenize a simple string", func() {
        s := Build()
        tokens := s.Tokenize("Hello world, to all the tests out there!")
        for _, str := range []string{"Hello", "world,", "to", "all", "the", "tests", "out", "there!"} {
            c.Expect((<-tokens).Backing(), Equals, str)
        }
    })
}
