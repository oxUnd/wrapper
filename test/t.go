package main

import (
    "fmt"
    "github.com/xiangshouding/wrapper"
)

func main() {
    w := wrapper.New(`<html> {{ob "test"}} function test() { {{.f}} {{.s}} {{.t}} } {{obE}} </html>`, func(t, c string) {
        fmt.Println(t, c)
    })

    w.Execute(map[string]interface{}{
        "f": "hello",
        "s": "fis",
        "t": "wrapper",
    })
}
