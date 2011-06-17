package stemmer

// #include "stemmer.h"
import "C"

import (
    "fetch/filter"
    "fetch/tokenizer"
)

type Stemmer struct {}

func (a *Stemmer) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    stemmer := C.create_stemmer()
    return filter.BuildFilter(in, func (t *tokenizer.Token) *tokenizer.Token {
        s := t.Backing()
        end := C.stem(stemmer, C.CString(s), C.int(len(s) - 1)) + 1
        return tokenizer.NewToken(s[0:end])
    })
}

func Build() *Stemmer {
    return new(Stemmer)
}
