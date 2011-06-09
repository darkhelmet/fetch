package tokenizer

type Tokenizer interface {
    Tokenize(input string) TokenChan
}
