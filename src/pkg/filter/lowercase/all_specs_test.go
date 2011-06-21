package lowercase

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(LowercaseSpec)
    gospec.MainGoTest(r, t)
}
