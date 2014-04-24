
package staticmap 


import (
    "bytes"
    "compress/gzip"
    "encoding/base64"
    "io"
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

func (self *Element) UnzipReader() io.Reader {
    buffer := bytes.NewBuffer(self.Bytes())
    reader, _ := gzip.NewReader(buffer)
    return reader
}

func (self *Element) UnzipBytes() []byte {
    buffer := new(bytes.Buffer)
    reader := self.UnzipReader()
    io.Copy(buffer, reader)
    return buffer.Bytes()
}

func (self *Element) AsString() string {
    return string(self.UnzipBytes())
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
    
    "hello.js": &Element{
        Rel : "hello.js",
        Ext : ".js",
        MimeType :"application/javascript" ,
        Data :"H4sIAAAJbogA/0rOzyvOz0nVy8lP11DKSM3JydfLKlbS5AIEAAD//3x2UNUYAAAA",
    },
    
}

