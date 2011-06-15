package fetch

import (
    "fetch/filter"
    "fetch/filter/ascii"
    "fetch/filter/punctuation"
    "fetch/filter/lowercase"
    "fetch/filter/superstrip"
    "fetch/filter/stopword"
    "fetch/tokenizer"
    "fetch/tokenizer/simple"
    "fetch/storage"
    "fetch/storage/redis"
)

type Engine struct {
    storage storage.Engine
}

func (e *Engine) Index(index, scope, id string, doc map[string]interface{}) {
    st := simple.Build()
    for field, v := range(doc) {
        text := v.(string)
        start, end := buildFilterChain("superstrip")
        go func() {
           for it := range(st.Tokenize(text)) {
               start <- it
           }
           close(start)
        }()
        go e.storage.Store(index, scope, id, field, end)
    }
}

func (e *Engine) Search(query string) []string {
    // TODO: Implement
    d := []string{query}
    return d
}

func buildFilterFromName(name string) filter.Filter {
    switch name {
    case "ascii": return ascii.Build()
    case "punctuation": return punctuation.Build()
    case "lowercase": return lowercase.Build()
    case "superstrip": return superstrip.Build()
    }
    panic("Invalid filter")
}

func buildFilterChain(names... string) (in, out tokenizer.TokenChan) {
    in = make(tokenizer.TokenChan, 10)
    head := buildFilterFromName(names[0])
    out = head.Process(in)
    for _, name := range(names[1:]) {
        out = buildFilterFromName(name).Process(out)
    }
    return in, out
}

func Build(engine string) *Engine {
    // TODO: Switch on engine
    return &Engine{ storage: redis.Build() }
}
