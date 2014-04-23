package wrapper

import (
	"bytes"
	"html/template"
)

func New(s string, cb func(typ, content string)) *Wrapper {
	handle := &Wrapper{}
	handle.buffer = bytes.NewBufferString("")
	_old := 0
	handle.typ = ""

	buildins := map[string]interface{}{
		"buffer": func(args ...interface{}) string {
			if len(args) == 1 {
				typ := args[0].(string)
				handle.typ = typ
			} else {
				Error("Must given one param!")
			}

			//old pos
			_old = handle.GetPos()

			return ""
		},
		"bufferEnd": func(args ...interface{}) string {

			//new pos
			_new := handle.GetPos()

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
		Error("initial template failed: \n" + err.Error())
	}

	return handle
}

func Error(msg string) {
	panic(msg)
}

type Wrapper struct {
	tmpl   *template.Template
	buffer *bytes.Buffer
	pos    int
	typ    string
}

func (w Wrapper) GetPos() int {
	w.pos = w.buffer.Len()
	return w.pos
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
