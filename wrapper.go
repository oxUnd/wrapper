package wrapper

import (
    "bytes"
    "html/template"
    "log"
)

func New(s string, cb func(typ, content string)) *Wrapper {
    handle := &Wrapper{}
    handle.buffer = bytes.NewBufferString("")
    handle.pos = 0
    handle.typ = ""

    buildins := map[string]interface{}{
        "ob": func(args ...interface{}) string {
            if len(args) == 1 {
                typ := args[0].(string)
                handle.typ = typ
            } else {

            }
            handle.Pos()
            return ""
        },
        "obE": func(args ...interface{}) string {
            _old := handle.pos
            _new := handle.Pos()
            content := handle.Read(_old, _new)
            typ := handle.typ
            //reset
            handle.Truncate(_old)
            //callback
            cb(typ, content)
            return ""
        },
    }

    var err error

    handle.tmpl, err = template.New(`template`).Funcs(buildins).Parse(s)

    if err != nil {
        log.Println("init tmplate failed!")
    }

    return handle
}

type Wrapper struct {
    tmpl   *template.Template
    buffer *bytes.Buffer
    pos    int
    typ    string
}

func (w Wrapper) GetPos() int {
    return w.pos
}

func (w Wrapper) Pos() int {
    p := w.buffer.Len()
    w.pos = p
    return p
}

func (w Wrapper) Truncate(n int) {
    w.buffer.Truncate(n)
}

func (w Wrapper) Read(start, end int) string {
    byte_buf := w.buffer.Bytes()
    buf := bytes.NewBuffer(byte_buf[start:end])
    return buf.String()
}

func (w Wrapper) Execute(data interface{}) *bytes.Buffer {
    w.tmpl.Execute(w.buffer, data)
    return w.buffer
}
