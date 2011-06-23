package main

import (
    // "os"
    "flag"
    "fmt"
    "fetch/config"
    // "scanner"
    "strings"
    // "io/ioutil"
    // "fetch"
    "github.com/darkhelmet/web.go"
)

func printResults(ids chan string) {
    fmt.Println("results")
    for id := range ids {
        fmt.Println(id)
    }
}

var configPath string

func init() {
    flag.StringVar(&configPath, "config", "", "Path to the config file")
}

func main() {
    flag.Parse()
    if configPath == "" {
        panic("Must specify config file")
    }

    c := config.NewFromFile(configPath)

    web.Get("/", func() string {
        return "Hello, World!"
    })

    web.Get("/(.*)/(.*)/(.*)", func(index, scope, id string) string {
        return strings.Join([]string{index, scope, id}, "/")
    })
    web.Run(c.Listen)

    // se := fetch.Build("redis")
    // se.Index("blog", "posts", "1", map[string]interface{}{"body": "Rubber duck debugging is when you have a bug, and can’t yet see what the problem is. You get a rubber duck, put it next to your monitor, and explain the problem you’re having. In the process of explaining the bug, you realize the root problem and are able to fix it."})
    // se.Index("blog", "posts", "2", map[string]interface{}{"body": "I also like to do debugging with my laptop."})
    // se.Index("blog", "pages", "3", map[string]interface{}{"body": "Seriously you guys, debugging is the best thing ever."})
    // printResults(se.SearchScope("blog", "posts", "debugging"))
    // printResults(se.SearchScope("blog", "posts", "laptop"))
    // printResults(se.SearchScope("blog", "pages", "debugging"))
}
