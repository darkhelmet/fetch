package main

import (
    // "os"
    "fmt"
    // "scanner"
    // "strings"
    // "io/ioutil"
    "fetch"
)

func printResults(ids chan string) {
    fmt.Println("results")
    for id := range(ids) {
        fmt.Println(id)
    }
}

func main() {
    // path := "/Users/darkhelmet/dev/github/darkhelmet/fetch/src/cmd/fetchd/fetchd.go"
    // b, _ := ioutil.ReadFile(path)
    // f := string(b)

    // file, _ := os.Open(path)
    // var s scanner.Scanner
    // s.Init(strings.NewReader("hello there what's going    on!"))
    // tok := s.Scan()
    // for tok != scanner.EOF {
    //     fmt.Println(s.TokenText())
    //     tok = s.Scan()
    // }

    se := fetch.Build("redis")
    se.Index("blog", "posts", "1", map[string]interface{}{"body": "Rubber duck debugging is when you have a bug, and can’t yet see what the problem is. You get a rubber duck, put it next to your monitor, and explain the problem you’re having. In the process of explaining the bug, you realize the root problem and are able to fix it."})
    se.Index("blog", "posts", "2", map[string]interface{}{"body": "I also like to do debugging with my laptop."})
    se.Index("blog", "pages", "3", map[string]interface{}{"body": "Seriously you guys, debugging is the best thing ever."})
    printResults(se.SearchScope("blog", "posts", "debugging"))
    printResults(se.SearchScope("blog", "posts", "laptop"))
    printResults(se.SearchScope("blog", "pages", "debugging"))
}
