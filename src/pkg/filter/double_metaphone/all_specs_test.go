package double_metaphone

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(DoubleMetaphoneSpec)
    gospec.MainGoTest(r, t)
}
