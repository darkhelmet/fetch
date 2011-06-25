package storage

import (
    "fetch/tokenizer"
)

type Engine interface {
    Store(index, scope, id, field string, tc tokenizer.TokenChan) bool
    Delete(index, scope, id string) bool
    SearchField(index, scope, field string, tc tokenizer.TokenChan) chan string
    SearchScope(index, scope string, tc tokenizer.TokenChan) chan string
}
