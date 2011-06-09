package fetch

import (
    "fmt"
    "fetch/storage"
    "fetch/storage/redis"
)

type Engine struct {
    storage storage.Engine
}

func (e *Engine) Index(index, scope, id string, doc map[string]interface{}) {
    for k, v := range(doc) {
        // TODO: Implement
        text := v.(string)
        fmt.Println(k)
        fmt.Println(text)
    }
}

func (e *Engine) Search(query string) []string {
    // TODO: Implement
    d := []string{query}
    return d
}

func Build(engine string) *Engine {
    // TODO: Switch on engine
    return &Engine{ storage: redis.Build() }
}
