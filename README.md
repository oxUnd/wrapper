A html/template Wrapper
======================

### DESC

```html
{{ob "script"}}
 //balabala...
{{obE}}
```

获取ob, obE中间夹着的内容, 并传给回掉函数，进行进一步的处理。

### INSTALL

```bash
$ go get github.com/xiangshouding/wrapper

```

### USE

```go
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
```

