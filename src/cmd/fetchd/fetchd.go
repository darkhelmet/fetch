package main

import (
    "flag"
    "fmt"
    "fetch"
    "fetch/config"
    "fetch/web"
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
    e := fetch.Build("redis")
    web.Start(e, &c)
}
