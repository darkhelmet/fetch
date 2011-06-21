package redis

import (
    "gospec"
    "testing"
)

func TestAllSpecs(t *testing.T) {
    r := gospec.NewRunner()
    r.AddSpec(RedisSpec)
    gospec.MainGoTest(r, t)
}
