package stemmer

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(StemmerSpec)
    gospec.MainGoTest(r, t)
}
