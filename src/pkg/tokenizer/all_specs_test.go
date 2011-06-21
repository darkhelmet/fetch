package tokenizer

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(TokenizerSpec)
    gospec.MainGoTest(r, t)
}
