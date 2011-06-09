package storage

import (
    "fetch/tokenizer"
)

type Engine interface {
    Store(index, scope, id, field string, tc tokenizer.TokenChan)
}
