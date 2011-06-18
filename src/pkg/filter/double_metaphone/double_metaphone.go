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

type DoubleMetaphone struct {}

// FIXME: Not so many mallocs, plz
func (dm *DoubleMetaphone) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    return filter.StartFilter(in, func(t *tokenizer.Token) []*tokenizer.Token {
        cs := C.CString(t.Backing())
        defer C.free(unsafe.Pointer(cs))
        codes := C.double_metaphone(cs)
        primary, secondary := C.GoString(codes.primary), C.GoString(codes.secondary)
        defer C.free_dm_result(codes)
        if primary == secondary {
            return []*tokenizer.Token{tokenizer.NewToken(primary)}
        }
        return []*tokenizer.Token{tokenizer.NewToken(primary), tokenizer.NewToken(secondary)}
    })
}

func (dm *DoubleMetaphone) Cleanup() {
}

func Build() *DoubleMetaphone {
    return new(DoubleMetaphone)
}
