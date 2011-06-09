package redis

import (
    "fetch/tokenizer"
)

type Redis struct {}

func (r *Redis) Store(index, scope, id, field string, tc tokenizer.TokenChan) {
    // TODO: Implement
}

func Build() (r *Redis) {
    // TODO: Configuration
    return &Redis{}
}
