package simple

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(SimpleSpec)
    gospec.MainGoTest(r, t)
}
