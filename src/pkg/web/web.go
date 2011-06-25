package web

import (
    "container/vector"
    "fmt"
    "json"
    "fetch"
    "fetch/config"
    "github.com/darkhelmet/web.go"
)

const documentRoute = "/(.*)/(.*)/(.*)"

type JSON map[string]interface{}

func quickJson(key, message string) []byte {
    raw, _ := json.Marshal(map[string]interface{}{
        key: message,
    })
    return raw
}

func resultsJson(results chan string) []byte {
    var ids vector.StringVector
    for id := range results {
        ids.Push(id)
    }
    raw, _ := json.Marshal(map[string]interface{}{
        "results": ids,
    })
    return raw
}

func Start(engine *fetch.Engine, c *config.Config) {
    web.Get("/", func() string {
        return "Go Fetch!"
    })

    // Search!
    web.Get("/(.*)/(.*)", func(ctx *web.Context, index, scope string) {
        ctx.SetHeader("Content-Type", "application/json; charset=utf-8", true)
        if query := ctx.Params["query"]; query == "" {
            ctx.StartResponse(400)
            ctx.Write(quickJson("error", "No query provided"))
        } else {
            ctx.Write(resultsJson(engine.SearchScope(index, scope, query)))
        }
    })

    // Add a document to the index, overwriting what was there.
    web.Post(documentRoute, func(ctx *web.Context, index, scope, id string) {
        ctx.SetHeader("Content-Type", "application/json; charset=utf-8", true)
        var document JSON
        if err := json.Unmarshal(ctx.ParamData, &document); err != nil {
            ctx.StartResponse(400)
            ctx.Write(quickJson("error", fmt.Sprintf("Error parsing JSON: %s", err.String())))
            return
        }
        engine.Index(index, scope, id, document)
        ctx.Write(quickJson("success", "ok"))
    })

    // Remove a document from the index
    web.Delete(documentRoute, func(ctx *web.Context, index, scope, id string) {
        engine.Delete(index, scope, id)
        ctx.SetHeader("Content-Type", "application/json; charset=utf-8", true)
        ctx.Write(quickJson("success", "ok"))
    })

    web.Run(c.Listen)
}
