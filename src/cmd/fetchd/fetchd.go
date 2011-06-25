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
var showVersion bool

func init() {
    flag.StringVar(&configPath, "config", "", "Path to the config file")
    flag.BoolVar(&showVersion, "version", false, "Show the version")
}

func main() {
    flag.Parse()
    versionString := fmt.Sprintf("fetch v%s © Daniel Huckstep", fetch.Version)

    if showVersion {
        fmt.Println(versionString)
        return
    }

    if configPath == "" {
        panic("Must specify config file")
    }

    fmt.Println("Starting", versionString)

    c := config.NewFromFile(configPath)
    e := fetch.Build("redis")
    web.Start(e, &c)
}
