package redis

import (
    "tideland-rdc.googlecode.com/hg"
    "fetch/tokenizer"
)

type Redis struct {
    redis *rdc.RedisDatabase
}

func (r *Redis) Store(index, scope, id, field string, tc tokenizer.TokenChan) {
    // TODO: Implement
}

func Build() (r *Redis) {
    return &Redis{ redis: rdc.NewRedisDatabase(rdc.Configuration{}) }
}
