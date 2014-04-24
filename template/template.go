package template

import (
    "text/template"
)

const TEMPLATE = `
package {{.Name}} 


import (
    "encoding/base64"
    "log"
    "net/http"
    "strconv"
)

type Element struct {
    Rel      string
    Ext      string
    MimeType string
    Data     string
}

func Get(path string) *Element {
    return smap[path]
}

func (self *Element) Bytes() []byte {
    return decodeBase64(self.Data)
}

func decodeBase64(input string) []byte {
    data, _ := base64.StdEncoding.DecodeString(input)
    return data
}

type StaticServer struct {
    Root string
}

func NewStaticServer(root string) (string, *StaticServer) {
    return root, &StaticServer{Root: root}
}
func (self *StaticServer) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[len(self.Root):]
    if element := Get(path); element != nil {
        log.Printf("%s", path)
        writer.Header().Set("Content-Type", element.MimeType)
        writer.Header().Set("Content-Encoding", "gzip")
        bytes := element.Bytes()
        writer.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
        writer.Write(bytes)
    } else {
        http.NotFound(writer, r)
    }
}


var smap = map[string]*Element{
    {{ range $item := .Items}}
    "{{$item.Rel}}": &Element{
        Rel : "{{$item.Rel}}",
        Ext : "{{$item.Ext}}",
        MimeType :"{{$item.MimeType}}" ,
        Data :"{{$item.Data}}",
    },
    {{end}}
}

`

var spackTemplate *template.Template

func GetTemplate() *template.Template {
    if spackTemplate != nil {
        return spackTemplate
    }
    funcMap := template.FuncMap{
        "every20": func(i int) bool {
            return (i % 20) == 0
        },
    }
    var err error
    spackTemplate, err = template.New("spack").Funcs(funcMap).Parse(TEMPLATE)
    if err != nil {
        panic(err)
    }
    return spackTemplate
}
