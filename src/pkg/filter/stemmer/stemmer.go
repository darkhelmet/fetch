package stemmer

// #include "stemmer.h"
import "C"

import (
    "unsafe"
    "fetch/filter"
    "fetch/tokenizer"
)

type Stemmer struct {}

func (s *Stemmer) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    cstemmer := C.create_stemmer()
    return filter.BuildFilter(in, func(t *tokenizer.Token) *tokenizer.Token {
        str := t.Backing()
        cs := C.CString(str)
        defer C.free(unsafe.Pointer(cs))
        end := C.stem(cstemmer, cs, C.int(len(str) - 1)) + 1
        return tokenizer.NewToken(str[0:end])
    }, func() {
        C.free_stemmer(cstemmer)
    })
}

func Build() *Stemmer {
    return new(Stemmer)
}
