package filter

import (
    "fetch/tokenizer"
)

type Filter interface {
    Process(in, out tokenizer.TokenChan)
}
