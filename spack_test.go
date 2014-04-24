package main

import (
    "bytes"
    "fmt"
    "github.com/hirokidaichi/go-static-gen/test/sample"
    "io"
    "testing"
)

var P = fmt.Println

func TestReader(t *testing.T) {
    reader, err := staticmap.Get("hello.js").Reader()
    if err != nil {
        t.Errorf("cannot create gzip reader :%s", err)
    }
    buffer := new(bytes.Buffer)
    io.Copy(buffer, reader)
    if content := string(buffer.Bytes()); content != "console.log(\"hello.js\")\n" {
        t.Errorf("does not match file content:\n%s\n%s", content)
    }
    if staticmap.Get("hello.js").Length != 24 {
        t.Errorf("expect 24,content length")
    }
}
