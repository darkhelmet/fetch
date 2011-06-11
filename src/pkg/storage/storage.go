package storage

import (
    "strings"
    "fetch/tokenizer"
)

type Engine interface {
    Store(index, scope, id, field string, tc tokenizer.TokenChan)
}

func BuildKey(index, scope, id, field string, token *tokenizer.Token) string {
    return strings.Join([]string{index, scope, id, field, token.Backing()}, ":")
}
