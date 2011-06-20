package punctuation

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(PunctuationSpec)
    gospec.MainGoTest(r, t)
}
