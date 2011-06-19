package double_metaphone

/*
#cgo CFLAGS: -Wno-parentheses
#include "double_metaphone.h"
*/
import "C"

import (
    "unsafe"
    "fetch/filter"
    "fetch/tokenizer"
)

type DoubleMetaphone struct{}

func (dm *DoubleMetaphone) Process(input tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(input, func(token *tokenizer.Token, output tokenizer.TokenChan) {
        cs := C.CString(token.Backing())
        defer C.free(unsafe.Pointer(cs))
        codes := C.double_metaphone(cs)
        primary, secondary := C.GoString(codes.primary), C.GoString(codes.secondary)
        defer C.free_dm_result(codes)
        output <- tokenizer.NewToken(primary)
        if primary != secondary {
            output <- tokenizer.NewToken(secondary)
        }
    })
}

func (dm *DoubleMetaphone) Cleanup() {
}

func Build() *DoubleMetaphone {
    return new(DoubleMetaphone)
}
