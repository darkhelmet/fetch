package fetch

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(FetchSpec)
    gospec.MainGoTest(r, t)
}
