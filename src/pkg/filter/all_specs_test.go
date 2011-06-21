package filter

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(FilterSpec)
    gospec.MainGoTest(r, t)
}
