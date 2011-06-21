package superstrip

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(SuperstripSpec)
    gospec.MainGoTest(r, t)
}
