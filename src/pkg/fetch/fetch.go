package fetch

import (
    "fetch/filter"
    "fetch/filter/ascii"
    "fetch/filter/punctuation"
    "fetch/filter/lowercase"
    "fetch/filter/superstrip"
    "fetch/filter/stopword"
    "fetch/filter/stemmer"
    "fetch/filter/double_metaphone"
    "fetch/tokenizer"
    "fetch/tokenizer/simple"
    "fetch/storage"
    "fetch/storage/redis"
)

type Engine struct {
    storage storage.Engine
}

type FilterChain struct {
    input, output tokenizer.TokenChan
    filters []filter.Filter
}

func (fc *FilterChain) Close() {
    close(fc.input)
    // TODO: Handle Cleanup() of filters
}

func (fc *FilterChain) Pump(tokens tokenizer.TokenChan) {
    for token := range(tokens) {
        fc.input <- token
    }
}

func buildChainAndTokenize(text string) tokenizer.TokenChan {
    st := simple.Build()
    chain := buildFilterChain("superstrip", "stopword", "stemmer", "double_metaphone")
    go func() {
       chain.Pump(st.Tokenize(text))
       chain.Close()
    }()
    return chain.output
}

func buildFilterChain(names... string) (*FilterChain) {
    var output tokenizer.TokenChan
    chain := &FilterChain{
        input: make(tokenizer.TokenChan, 10),
        filters: make([]filter.Filter, len(names)),
    }
    output = chain.input
    for index, name := range(names) {
        chain.filters[index] = buildFilterFromName(name)
        output = chain.filters[index].Process(output)
    }
    chain.output = output
    return chain
}

func buildFilterFromName(name string) filter.Filter {
    switch name {
    case "ascii": return ascii.Build()
    case "punctuation": return punctuation.Build()
    case "lowercase": return lowercase.Build()
    case "superstrip": return superstrip.Build()
    case "stopword": return stopword.Build()
    case "stemmer": return stemmer.Build()
    case "double_metaphone": return double_metaphone.Build()
    }
    panic("Invalid filter")
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
