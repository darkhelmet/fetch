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

func buildChainAndTokenize(text string) tokenizer.TokenChan {
    st := simple.Build()
    start, end := buildFilterChain("superstrip", "stopword")
    go func() {
       for it := range(st.Tokenize(text)) {
           start <- it
       }
       close(start)
    }()
    return end
}

func buildFilterFromName(name string) filter.Filter {
    switch name {
    case "ascii": return ascii.Build()
    case "punctuation": return punctuation.Build()
    case "lowercase": return lowercase.Build()
    case "superstrip": return superstrip.Build()
    case "stopword": return stopword.Build()
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

func (e *Engine) Index(index, scope, id string, doc map[string]interface{}) {
    for field, v := range(doc) {
        text := v.(string)
        end := buildChainAndTokenize(text)
        go e.storage.Store(index, scope, id, field, end)
    }
}

func (e *Engine) SearchField(index, scope, field, query string) chan string {
    return e.storage.SearchField(index, scope, field, buildChainAndTokenize(query))
}

func (e *Engine) SearchScope(index, scope, query string) chan string {
    return e.storage.SearchScope(index, scope, buildChainAndTokenize(query))
}

func Build(engine string) *Engine {
    // TODO: Switch on engine
    return &Engine{ storage: redis.Build() }
}
