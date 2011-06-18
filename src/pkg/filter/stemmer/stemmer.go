package stemmer

// #include "stemmer.h"
import "C"

import (
    "unsafe"
    "fetch/filter"
    "fetch/tokenizer"
)

type Stemmer struct {
    cstemmer *C.struct_stemmer
}

func (s *Stemmer) Process(input tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(input, func(token *tokenizer.Token, output tokenizer.TokenChan) {
        str := token.Backing()
        cs := C.CString(str)
        defer C.free(unsafe.Pointer(cs))
        end := C.stem(s.cstemmer, cs, C.int(len(str) - 1)) + 1
        output <- tokenizer.NewToken(str[0:end])
    })
}

func (s *Stemmer) Cleanup() {
    C.free_stemmer(s.cstemmer)
}

func Build() *Stemmer {
    return &Stemmer{ cstemmer: C.create_stemmer() }
}
