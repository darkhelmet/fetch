package tokenizer

type Token struct {
    // Backing string
    backing string
}

func (t *Token) Backing() string {
    return t.backing
}

type TokenChan chan *Token

// Make a token frm a string
func NewToken(from string) *Token {
    return &Token{backing: from}
}

func NewTokenChan(strs []string) TokenChan {
    tc := make(TokenChan, 10)
    go func() {
        for _, s := range strs {
            tc <- NewToken(s)
        }
        close(tc)
    }()
    return tc
}
