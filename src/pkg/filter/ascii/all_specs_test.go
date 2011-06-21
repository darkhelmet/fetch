package ascii

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(AsciiSpec)
    gospec.MainGoTest(r, t)
}
