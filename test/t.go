package main

import (
    "fmt"
    "github.com/xiangshouding/wrapper"
)

func main() {
    w := wrapper.New(`{{ob}}{{.f}} {{.s}} {{.t}}{{obE}}`, func(t, c string) {
        fmt.Println(t, c)
    })

    fmt.Println(w.Execute(map[string]interface{}{
        "f": "hello",
        "s": "fis",
        "t": "wrapper",
    }))
}
