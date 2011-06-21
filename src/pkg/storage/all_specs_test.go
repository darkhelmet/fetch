package storage

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(StorageSpec)
    gospec.MainGoTest(r, t)
}
