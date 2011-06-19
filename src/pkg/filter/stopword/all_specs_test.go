package stopword

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(StopwordSpec)
    gospec.MainGoTest(r, t)
}
