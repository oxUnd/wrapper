A html/template Wrapper
======================

### DESC

```html
{{buffer "script"}}
 //balabala...
{{bufferEnd}}
```

it like `ob_start(), ob_end()` of PHP, get buffer content;

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
    w := wrapper.New(`{{buffer "script"}}{{.f}} {{.s}} {{.t}}{{bufferEnd}}`, func(t, c string) {
        fmt.Println(t, c)
    })

    fmt.Println(w.Execute(map[string]interface{}{
        "f": "hello",
        "s": "fis",
        "t": "wrapper",
    }))
}
```
### API

#### wrapper.New

```go
New(tmplStr string, cb func(typ, content string))

// typ "script" | "style"
// content  "script" | "css style"
func cb (typ, content) {
    //blabla...
}
```