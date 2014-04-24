
package staticmap 

import (
    "bytes"
    "compress/gzip"
    "encoding/base64"
    "io"
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
    return decodeBase64(self.Data))
}

func decodeBase64(input string) []byte {
    data, _ := base64.StdEncoding.DecodeString(input)
    return data
}

var smap = map[string]*Element{
    
    "hello.js": &Element{
        Rel : "hello.js",
        Ext : ".js",
        MimeType :"application/javascript" ,
        Data :"H4sIAAAJbogA/0rOzyvOz0nVy8lP11DKSM3JydfLKlbS5AIEAAD//3x2UNUYAAAA",
    },
    
}

